package main

import (
	"os"

	"github.com/ElladanTasartir/golang-amqp-publisher/application"
	"github.com/ElladanTasartir/golang-amqp-publisher/config"
	"github.com/ElladanTasartir/golang-amqp-publisher/service"
)

func main() {
	config.LoadEnv()

	conn, channel, queue := application.ConnectAMQP(
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_GOLANG_QUEUE"),
	)

	defer channel.Close()
	defer conn.Close()

	service.ConsumeQueue(channel, queue)
}
