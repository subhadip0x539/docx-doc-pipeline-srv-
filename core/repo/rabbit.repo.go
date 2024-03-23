package repo

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type IRabbitRepo interface {
}

type RabbitRepo struct {
	channel  *amqp.Channel
	exchange string
}

func NewRabbitRepo(channel *amqp.Channel, exchange string) IRabbitRepo {
	return &RabbitRepo{channel: channel, exchange: exchange}
}
