package config

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func amqpConn() (*amqp091.Channel, amqp091.Queue) {
	conn, err := amqp091.Dial(viper.GetString("AMQP_URI"))
	if err != nil {
		panic(err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}

	q, err := ch.QueueDeclare(
		viper.GetString("AMQP_NAME"),
		false,
		false,
		false,
		false,
		nil,
	)

	return ch, q
}
