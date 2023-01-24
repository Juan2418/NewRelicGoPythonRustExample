package main

import (
	"log"
	"new-relic-example/logs"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"

	_ "github.com/joho/godotenv/autoload"
)

func LogExampleRequest(c *gin.Context) {
	var message Message
	c.BindJSON(&message)

	requestId := c.MustGet("requestId").(string)
	log.Default().Printf("Logging message from request %s: \n%s\n", requestId, message.Message)

	c.JSON(200, gin.H{
		"requestId": requestId,
		"message":   message.Message,
	})
}

func LogReferenceRequest(c *gin.Context) {
	requestId := c.MustGet("requestId").(string)

	returnMessage := "This comes from Go with the reference " + requestId
	log.Default().Printf(returnMessage)

	c.JSON(200, gin.H{
		"message":   returnMessage,
		"requestId": requestId,
	})
}

func main() {
	newRelicKey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("NewRelicExample"),
		newrelic.ConfigLicense(newRelicKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	engine := gin.Default()
	engine.Use(logs.Logger())
	engine.Use(nrgin.Middleware(app))

	engine.POST("/", LogExampleRequest)
	engine.POST("/reference", LogReferenceRequest)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	engine.Run(":" + PORT)
}
