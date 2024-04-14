package v1

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitMQ struct {
	ch       *amqp.Channel
	exchange string
}

func newRabbitMQ(ch *amqp.Channel, exchange string) *RabbitMQ {
	return &RabbitMQ{
		ch:       ch,
		exchange: exchange,
	}
}
