package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	uri     string
	timeout int64
	channel *amqp.Channel
}

type IRabbit interface {
	Connect() error
	Disconnect() error
	GetChannel() *amqp.Channel
}

func (r *Rabbit) Connect() error {
	conn, err := amqp.Dial(r.uri)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	r.channel = ch

	return nil
}

func (r *Rabbit) Disconnect() error {
	if err := r.channel.Close(); err != nil {
		return err
	}

	return nil
}

func (r *Rabbit) GetChannel() *amqp.Channel {
	return r.channel
}

func NewRabbit(uri string, timeout int64) IRabbit {
	return &Rabbit{uri: uri, timeout: timeout}
}
