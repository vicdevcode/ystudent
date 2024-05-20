package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/internal/usecase"
)

type moderatorRoute struct {
	rmq *RabbitMQ
	u   usecase.User
	l   *slog.Logger
}

func newModerator(router *router, rmq *RabbitMQ, u usecase.User, l *slog.Logger) {
	r := &moderatorRoute{rmq, u, l}
	{
		router.protected.POST("/moderator/", r.create)
		router.protected.GET("/moderators/", r.findAll)
	}
}

type createModeratorRequest dto.CreateUser

type createModeratorResponse *entity.User

func (r *moderatorRoute) create(c *gin.Context) {
	var body createModeratorRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	moderator, err := r.u.Create(c.Request.Context(), dto.CreateUser{
		Fio:      body.Fio,
		Email:    body.Email,
		RoleType: entity.MODERATOR,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(moderator)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createModeratorResponse(moderator))

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.moderator.created",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}

type findAllModeratorResponse struct {
	Moderators []entity.User `json:"moderators"`
}

func (r *moderatorRoute) findAll(c *gin.Context) {
	moderators, err := r.u.FindAll(c.Request.Context(), entity.MODERATOR)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllModeratorResponse{Moderators: moderators})
}
