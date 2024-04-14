package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type groupRoute struct {
	u   usecase.Group
	l   *slog.Logger
	rmq *RabbitMQ
}

func newGroup(handler *gin.RouterGroup, u usecase.Group, rmq *RabbitMQ, l *slog.Logger) {
	r := &groupRoute{u, l, rmq}
	h := handler.Group("/group")
	{
		h.POST("/", r.create)
		h.GET("/", r.findAll)
	}
}

type createGroupResponse struct {
	*entity.Group
	dto.CUD
}

func (r *groupRoute) create(c *gin.Context) {
	var body dto.CreateGroup

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	group, err := r.u.Create(c.Request.Context(), body)
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createGroupResponse{Group: group})
}

type findAllGroupResponse struct {
	Groups []entity.Group `json:"groups"`
}

func (r *groupRoute) findAll(c *gin.Context) {
	groups, err := r.u.FindAll(c.Request.Context())
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(groups)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, findAllGroupResponse{Groups: groups})

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"lol.groups",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}
