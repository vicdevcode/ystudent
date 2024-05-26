package v1

import (
	"encoding/json"
	"fmt"
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
	ug  usecase.Group
	ud  usecase.Department
	ut  usecase.Teacher
	l   *slog.Logger
}

func newGroup(
	router *router,
	rmq *RabbitMQ,
	ug usecase.Group,
	ud usecase.Department,
	ut usecase.Teacher,
	l *slog.Logger,
) {
	r := &groupRoute{rmq, ug, ud, ut, l}
	{
		router.protected.POST("/group/", r.create)
		router.protected.PUT("/group/:id", r.update)
		router.protected.DELETE("/group/:id", r.delete)
		// router.protected.PUT("/group/update_curator/:id", r.updateCurator)
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

	group, err := r.ug.Create(c.Request.Context(), body)
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
	Groups []dto.FindAllGroupResponse `json:"groups"`
	dto.Page
}

func (r *groupRoute) findAll(c *gin.Context) {
	page, err := GetPage(c.Query("page"), c.Query("count"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	groups, err := r.ug.FindAll(c.Request.Context(), page)
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	var response []dto.FindAllGroupResponse
	for _, group := range groups {
		department, err := r.ud.FindOne(c.Request.Context(), entity.Department{
			ID: *group.DepartmentID,
		})
		if err != nil {
			r.l.Error(err.Error())
			internalServerError(c, err.Error())
			return
		}
		var curatorFio string
		if group.CuratorID != nil {
			curator, err := r.ut.FindOne(c.Request.Context(), entity.Teacher{
				ID: *group.CuratorID,
			})
			if err != nil {
				r.l.Error(err.Error())
				internalServerError(c, err.Error())
				return
			}
			if curator.User.Middlename != "" {
				curatorFio = fmt.Sprintf(
					"%s %s %s",
					curator.User.Surname,
					curator.User.Firstname,
					curator.User.Middlename,
				)
			} else {
				curatorFio = fmt.Sprintf("%s %s", curator.User.Surname, curator.User.Firstname)
			}
		}
		response = append(response, dto.FindAllGroupResponse{
			Group:          &group,
			DepartmentName: department.Name,
			CuratorFio:     curatorFio,
		})
	}

	c.JSON(http.StatusOK, findAllGroupResponse{Groups: response, Page: page})
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

	group, err := r.ug.UpdateCurator(c.Request.Context(), dto.UpdateGroupCurator{
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

type updateGroupRequest dto.UpdateGroupBody

type updateGroupResponse *entity.Group

func (r *groupRoute) update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	var body updateGroupRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	data := dto.UpdateGroup{
		ID: id,
	}
	if body.Name != "" {
		data.Name = body.Name
	}
	if body.CuratorID != uuid.Nil {
		data.CuratorID = &body.CuratorID
	}
	if body.DepartmentID != uuid.Nil {
		data.DepartmentID = &body.DepartmentID
	}

	group, err := r.ug.Update(c.Request.Context(), data)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateGroupResponse(group))
}

type deleteGroupResponse struct {
	Message string `json:"message"`
}

func (r *groupRoute) delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	if err := r.ug.Delete(c.Request.Context(), id); err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, deleteGroupResponse{
		Message: "group was deleted",
	})
}
