package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Client for sending requests on behalf of bot
type Client struct {
	token string
}

// GetAPIBaseURL returns telegram base URL for sending requests
func (c *Client) GetAPIBaseURL() string {
	return "https://api.telegram.org/"
}

func (c *Client) GetAPIUrl(method string) string {
	return fmt.Sprintf("%sbot%s/%s", c.GetAPIBaseURL(), c.token, method)
}

// NewClient Creates new telegram client for sending requests
func NewClient(botToken string) *Client {
	return &Client{
		token: botToken,
	}
}

// GetWebhookInfo gets webhook current data
func (c *Client) GetWebhookInfo() (WebhookInfo, error) {
	var info WebhookInfo
	resp, err := http.Get(c.GetAPIUrl("getWebhookInfo"))
	if err != nil {
		return info, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &info)
	if err != nil {
		return info, err
	}

	return info, nil
}

// SetWebhook sets webhook URL
func (c *Client) SetWebhook(url string, param SetWebhookParams) error {
	requestData := map[string]string{
		"url": url,
	}
	if param.PublicCertificateURL != "" {
		requestData["certificate"] = param.PublicCertificateURL
	}
	if len(param.AllowedUpdates) > 0 {
		str, err := json.Marshal(param.AllowedUpdates)
		if err != nil {
			return err
		}
		requestData["allowed_updates"] = string(str)
	}
	if param.MaxConnections > 0 {
		requestData["max_connections"] = strconv.Itoa(param.MaxConnections)
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	requestBodyReader := bytes.NewReader(requestBody)

	resp, err := http.Post(c.GetAPIUrl("setWebhook"), "application/json", requestBodyReader)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(fmt.Sprintf("Could not set webhook: %s", string(bodyBytes)))
	}

	return nil
}
