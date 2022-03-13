package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	// URL for webhooks from telegram. Should contain only base part without trailing slash
	WebhookURL string
	// Host for webserver
	Host string
	// Port for webserver
	Port int
	// Secret key for validating that message is from telegram
	WebhookSecretKey string
	// Telegram bot token for sending requests
	TelegramBotToken string
}

// LoadConfig loads application configuration from file
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileBody, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(fileBody, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig saves application configuration to file
func (cfg *Config) SaveConfig(filename string) error {
	bytes, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, bytes, 0666)

	return err
}
