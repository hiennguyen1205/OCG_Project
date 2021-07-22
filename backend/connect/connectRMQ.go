package connect

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Conn struct {
	Channel *amqp.Channel
}

func PublishingAMessage(data []byte) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Println("Successfuly Connected To our RMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"EmailQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"EmailQueue",
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
}
