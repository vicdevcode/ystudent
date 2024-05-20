package consumer

import (
	"context"
	"log/slog"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/vicdevcode/ystudent/main/internal/usecase"
)

type Consumer struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	queue    amqp.Queue
	ctx      context.Context
	delivery <-chan amqp.Delivery
}

func New(
	conn *amqp.Connection,
	ch *amqp.Channel,
	ctx context.Context,
	delivery <-chan amqp.Delivery,
) *Consumer {
	return &Consumer{
		conn: conn, ch: ch,
		ctx:      ctx,
		delivery: delivery,
	}
}

func (c *Consumer) Start(uc usecase.UseCases, l *slog.Logger) {
	for {
		select {
		case <-c.ctx.Done():
			return
		case d, opened := <-c.delivery:
			if !opened {
				l.Error("Consumer closed")
				return
			}
			switch d.RoutingKey {
			// case "lol.faculties":
			// 	var response []entity.Faculty
			// 	if err := json.Unmarshal(d.Body, &response); err != nil {
			// 		l.Error(err.Error())
			// 	} else {
			// 		l.Info(fmt.Sprintf("%v", response))
			// 	}
			// 	l.Debug(d.RoutingKey)
			// 	break
			// case "lol.groups":
			// 	var response []entity.Group
			// 	if err := json.Unmarshal(d.Body, &response); err != nil {
			// 		l.Error(err.Error())
			// 		break
			// 	}
			// 	student, err := uc.UserUseCase.FindOne(context.Background(), entity.User{
			// 		ID: response[0].Students[0].UserID,
			// 	})
			// 	if err != nil {
			// 		l.Error(err.Error())
			// 		break
			// 	}
			// 	l.Info(student.Firstname)
			// 	l.Debug(d.RoutingKey)
			// 	break
			default:
				l.Info(d.RoutingKey)
				break
			}
			d.Ack(false)
		}
	}
}
