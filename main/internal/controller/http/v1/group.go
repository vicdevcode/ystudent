package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/internal/usecase"
)

type groupRoute struct {
	rmq *RabbitMQ
	u   usecase.Group
	l   *slog.Logger
}

func newGroup(
	router *router,
	rmq *RabbitMQ,
	u usecase.Group,
	l *slog.Logger,
) {
	r := &groupRoute{rmq, u, l}
	{
		router.protected.POST("/group/", r.create)
		router.protected.PUT("/group/:id", r.updateCurator)
		router.public.GET("/groups/", r.findAll)
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

	response, err := json.Marshal(group)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, createGroupResponse{Group: group})

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.group.created",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
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

	c.JSON(http.StatusOK, findAllGroupResponse{Groups: groups})
}

type updateCurator struct {
	TeacherId uuid.UUID `json:"teacher_id"`
}

func (r *groupRoute) updateCurator(c *gin.Context) {
	var body updateCurator

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	group, err := r.u.UpdateCurator(c.Request.Context(), dto.UpdateGroupCurator{
		ID:        id,
		CuratorID: body.TeacherId,
	})
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createGroupResponse{Group: group})

	response, err := json.Marshal(group)
	if err != nil {
		return
	}

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.group.curator_updated",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}
