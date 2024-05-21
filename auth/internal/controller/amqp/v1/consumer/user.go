package consumer

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type userRoute struct {
	uh usecase.Hash
	uu usecase.User
	l  *slog.Logger
}

func newUser(
	uh usecase.Hash,
	uu usecase.User,
	l *slog.Logger,
) *userRoute {
	return &userRoute{uh, uu, l}
}

func (r *userRoute) created(c *Consumer, d amqp091.Delivery) error {
	var user *entity.User
	if err := json.Unmarshal(d.Body, &user); err != nil {
		return err
	}

	hashedPassword, err := r.uh.HashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = r.uu.Create(c.ctx, dto.CreateUser{
		ID: user.ID,
		Fio: dto.Fio{
			Firstname:  user.Firstname,
			Middlename: user.Middlename,
			Surname:    user.Surname,
		},
		Email:    user.Email,
		Password: hashedPassword,
		Role:     entity.UserRole(user.RoleType),
	})
	if err != nil {
		return err
	}

	r.l.Info(fmt.Sprintf("%s was created", user.RoleType))
	d.Acknowledger.Ack(d.DeliveryTag, false)
	return nil
}
