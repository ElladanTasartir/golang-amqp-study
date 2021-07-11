package main

import (
	"os"

	"github.com/ElladanTasartir/golang-amqp-publisher/application"
	"github.com/ElladanTasartir/golang-amqp-publisher/config"
)

func main() {
	config.LoadEnv()

	amqpClient := application.ConnectAMQP(
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_GOLANG_QUEUE"),
	)

	router := application.StartHTTP()

	defer amqpClient.Channel.Close()
	defer amqpClient.Conn.Close()

	router.Run(":8080")
}
