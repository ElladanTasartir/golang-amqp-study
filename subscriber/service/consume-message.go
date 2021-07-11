package service

import (
	"fmt"

	"github.com/ElladanTasartir/golang-amqp-publisher/util"
	"github.com/streadway/amqp"
)

func ConsumeQueue(channel *amqp.Channel, queue *amqp.Queue) {
	messages, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		util.FailOnError(err, "Failed to consume")
	}

	go readFromQueue(messages)

	fmt.Println(
		fmt.Sprintf("Messages are being read from queue %s", queue.Name),
	)
}

func readFromQueue(messages <-chan amqp.Delivery) {
	for message := range messages {
		fmt.Println(
			fmt.Sprintf("Received message: %s", message.Body),
		)
	}
}
