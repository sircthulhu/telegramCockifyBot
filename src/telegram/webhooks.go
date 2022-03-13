package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"telegramCockifyBot/src/util"
)

// HandleWebhook handles webhooks from telegram
func HandleWebhook(tgClient Client, c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error while reading body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not read request body"})
		return
	}
	upd := update{}
	err = json.Unmarshal(bodyBytes, &upd)
	if err != nil {
		log.Printf("Error while parsing body JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request body JSON"})
		return
	}

	if upd.InlineQuery.ID != "" {
		err = handleInlineQuery(upd, tgClient)
		if err != nil {
			log.Printf("Error while answering inline query: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while answering inline query"})
		}
	}

	c.String(http.StatusOK, "")
}

// handleInlineQuery handles new inline query
func handleInlineQuery(update update, tgClient Client) error {
	answer := answerInlineQueryParams{
		InlineQueryID: update.InlineQuery.ID,
		CacheTime:     60, //int(time.Hour * 12),
		IsPersonal:    true,
		Results: []inlineQueryResultArticle{
			{
				Type:        "article",
				ID:          "1",
				Title:       "How long today?",
				Description: "Share your cock size",
				InputMessageContent: inputMessageContent{
					MessageText: util.FormatCockSizeMessage(util.GenerateCockSize()),
				},
			},
		},
	}
	encodedAnswer, err := json.Marshal(answer)
	if err != nil {
		return err
	}
	answerReader := bytes.NewReader(encodedAnswer)
	resp, err := http.Post(tgClient.GetAPIUrl("answerInlineQuery"), "application/json", answerReader)
	if err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Error while sending answerInlineQuery: %s", string(bodyBytes)))
	}

	return nil
}
