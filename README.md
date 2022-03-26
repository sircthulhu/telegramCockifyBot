# Telegram Cockify bot

Inline bot that sends random number (your cock size) to the chat  
Example message: "My cock size is 24 cm."

## Installation instruction
1. Install golang version 1.17
2. `go mod download && go mod tidy && go mod vendor`
3. MacOS M1 instruction for compilation
   1. Compile for linux `GOOS=linix GOARCH=amd64 go build -v src/main.go`
   2. Compile for MacOS `go build -v src/main.go`
4. Then you will have `main` file in the root directory of this project

## Starting
Program requires path to the config file as first argument (if this file does not exist, it will be created)  
Example starting command: `./main config.json` (this will create config.json file in this directory)

If config file does not exist, program will create empty config and exit. You should fill it.  
At start program sets webhook URL for telegram and starts waiting for connections

## Configuration
Example configuration:  
```
{
    "WebhookURL": "https url", // external URL for bot. It will be used for telegram webhook URL
    "Host": "127.0.0.1", // host on what webserver will be listening
    "Port": 8181, // port that webserver will be listening to
    "WebhookSecretKey": "secret_key", // secret key, so that we can verify that connections is from telegram
    "TelegramBotToken": "bot_token" // This is telegram bot token, can be acquired by botfather
}
```
For testing purpose, you can use `ngrok` and set WebhookURL to your ngrok https URL

## In production
There is `cockify.service` file in the root of this project that you can use for installing systemd service of bot.  
Important: set environment variable `GIN_MODE=release` in production