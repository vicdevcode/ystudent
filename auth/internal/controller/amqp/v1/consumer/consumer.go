package consumer

import (
	"context"
	"log/slog"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/vicdevcode/ystudent/auth/internal/usecase"
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
	adminRoute := newAdmin(uc.HashUseCase, uc.UserUseCase, l)
	userRoute := newUser(uc.HashUseCase, uc.UserUseCase, l)
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
			case "main.admin.created":
				if err := adminRoute.created(c, d); err != nil {
					l.Error(err.Error())
				}
				break
			case "main.student.created":
				if err := userRoute.created(c, d); err != nil {
					l.Error(err.Error())
				}
				break
			case "main.teacher.created":
				if err := userRoute.created(c, d); err != nil {
					l.Error(err.Error())
				}
				break
			case "main.employee.created":
				if err := userRoute.created(c, d); err != nil {
					l.Error(err.Error())
				}
				break
			case "main.moderator.created":
				if err := userRoute.created(c, d); err != nil {
					l.Error(err.Error())
				}
				break
			default:
				l.Info(d.RoutingKey)
				d.Acknowledger.Ack(d.DeliveryTag, false)
				break
			}
		}
	}
}
