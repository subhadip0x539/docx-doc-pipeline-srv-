package repo

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type IRabbitRepo interface {
}

type RabbitRepo struct {
	channel *amqp.Channel
}

func NewRabbitRepo(channel *amqp.Channel) IRabbitRepo {
	return &RabbitRepo{channel: channel}
}
