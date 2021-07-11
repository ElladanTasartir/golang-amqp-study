package application

import (
	"fmt"
	"net/http"

	"github.com/ElladanTasartir/golang-amqp-publisher/service"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func StartHTTP() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	declareRoutes(router)

	return router
}

func declareRoutes(router *gin.Engine) {
	router.POST("/new-message", postNewMessage)
}

func postNewMessage(context *gin.Context) {
	var message Message
	context.BindJSON(&message)
	service.PublishMessage(AmqpClient.Channel, AmqpClient.Queue, message.Message)
	context.String(http.StatusCreated, fmt.Sprintf("Message %s was sent", message.Message))
}
