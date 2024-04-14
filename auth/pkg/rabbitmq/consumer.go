package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	queue    amqp.Queue
	ctx      context.Context
	delivery <-chan amqp.Delivery
}

func NewConsumer(
	conn *amqp.Connection,
	ch *amqp.Channel,
	ctx context.Context,
	delivery <-chan amqp.Delivery,
) *Consumer {
	return &Consumer{
		conn:     conn,
		ch:       ch,
		ctx:      ctx,
		delivery: delivery,
	}
}

func (c *Consumer) Start() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case d, opened := <-c.delivery:
			if !opened {
				fmt.Println("Consumer closed")
				return
			}
			fmt.Println("Received message:", d.Body)
			d.Ack(false)
		}
	}
}
