package service

import (
	"fmt"

	"github.com/ElladanTasartir/golang-amqp-publisher/util"
	"github.com/streadway/amqp"
)

func PublishMessage(channel *amqp.Channel, queue *amqp.Queue, body string) {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}

	err := channel.Publish(
		"",
		queue.Name,
		false,
		false,
		message,
	)

	if err != nil {
		util.FailOnError(err, "Failed to send the message")
	}

	fmt.Println(
		fmt.Sprintf("Message sent to queue %s", queue.Name),
	)
}
