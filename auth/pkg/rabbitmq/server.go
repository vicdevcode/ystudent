package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	ExchangeName string `yaml:"exchange_name"`
	QueueName    string `yaml:"queue_name"`
	Topic        string `yaml:"topic"`
}

func New(cfg *Config) (*amqp.Connection, *amqp.Channel, amqp.Queue, <-chan amqp.Delivery) {
	conn, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port),
	)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	err = ch.ExchangeDeclare(
		cfg.ExchangeName,
		"topic",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	queue, err := ch.QueueDeclare(
		cfg.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = ch.QueueBind(queue.Name, cfg.Topic, cfg.ExchangeName, false, nil)
	if err != nil {
		panic(err)
	}

	delivery, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	return conn, ch, queue, delivery
}
