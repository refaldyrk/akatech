package config

import (
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Config struct {
	DB          *gorm.DB
	Q           amqp091.Queue
	AMQPChannel *amqp091.Channel
}

func NewConfig() *Config {
	config := Config{}

	config.DB = connectPostgres()

	amqp, q := amqpConn()

	config.AMQPChannel = amqp
	config.Q = q

	return &config
}
