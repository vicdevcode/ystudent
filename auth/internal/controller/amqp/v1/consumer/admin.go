package consumer

import (
	"encoding/json"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type adminRoute struct {
	uh usecase.Hash
	uu usecase.User
	l  *slog.Logger
}

func newAdmin(
	uh usecase.Hash,
	uu usecase.User,
	l *slog.Logger,
) *adminRoute {
	return &adminRoute{uh, uu, l}
}

func (r *adminRoute) created(c *Consumer, d amqp091.Delivery) error {
	var admin *entity.User
	if err := json.Unmarshal(d.Body, &admin); err != nil {
		return err
	}

	hashedPassword, err := r.uh.HashPassword("123123123")
	if err != nil {
		return err
	}

	_, err = r.uu.Create(c.ctx, dto.CreateUser{
		ID:       admin.ID,
		Email:    admin.Email,
		Password: hashedPassword,
		Role:     entity.UserRole(admin.RoleType),
	})
	if err != nil {
		return err
	}

	r.l.Info("admin was created")
	d.Acknowledger.Ack(d.DeliveryTag, false)
	return nil
}
