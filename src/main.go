package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"telegramCockifyBot/src/telegram"
	"telegramCockifyBot/src/util"
)

var Configuration *util.Config
var tgClient *telegram.Client

func getWebhooksUrl() string {
	return Configuration.WebhookURL + "/webhook/" + Configuration.WebhookSecretKey
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/webhook/:key", func(c *gin.Context) {
		key := c.Param("key")
		if key != Configuration.WebhookSecretKey {
			c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
			return
		}

		telegram.HandleWebhook(*tgClient, c)
	})

	return r
}

// setupTelegramWebhooks sets up telegram webhooks
func setupTelegramWebhooks(c *telegram.Client) {
	info, err := c.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	webhookURL := getWebhooksUrl()
	if info.URL != webhookURL {
		log.Println("Current telegram webhooks url is different from the url given in configuration.")
		err = c.SetWebhook(webhookURL, telegram.SetWebhookParams{
			AllowedUpdates: []string{"inline_query"},
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must pass the path to config json file as first argument")
	}
	var err error

	Configuration, err = util.LoadConfig(os.Args[1])
	if err != nil {
		log.Println("Configuration could not be read from file.")
		if os.IsNotExist(err) {
			Configuration = &util.Config{}
			_ = Configuration.SaveConfig(os.Args[1])
			log.Fatalf("Empty configuration file has been created at %s\n", os.Args[1])
		}
		return
	}
	tgClient = telegram.NewClient(Configuration.TelegramBotToken)
	setupTelegramWebhooks(tgClient)
	log.Println("Telegram webhooks are set up")

	router := setupRouter()
	address := Configuration.Host + ":" + strconv.Itoa(Configuration.Port)
	gin.SetMode(gin.ReleaseMode)
	err = router.Run(address)
	if err != nil {
		log.Fatalf("Could not start webserver on %s: %v", address, err)
		return
	}
}
