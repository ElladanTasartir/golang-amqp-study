package application

import (
	"fmt"

	"github.com/ElladanTasartir/golang-amqp-publisher/util"
	"github.com/streadway/amqp"
)

type AMQPClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

var AmqpClient *AMQPClient

func ConnectAMQP(host string, queueName string) *AMQPClient {
	conn, err := amqp.Dial(host)
	if err != nil {
		util.FailOnError(err, "Failed to connect to RabbitMQ")
	}

	channel, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")

	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		util.FailOnError(err, "Failed to declare queue")
	}

	fmt.Println(
		fmt.Sprintf("Successfully started RabbitMQ %s", host),
	)

	AmqpClient = &AMQPClient{
		Conn:    conn,
		Channel: channel,
		Queue:   &queue,
	}

	return AmqpClient
}
