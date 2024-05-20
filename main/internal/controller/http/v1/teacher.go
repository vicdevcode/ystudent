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

type teacherRoute struct {
	rmq *RabbitMQ
	ut  usecase.Teacher
	uu  usecase.User
	l   *slog.Logger
}

func newTeacher(
	router *router,
	rmq *RabbitMQ,
	ut usecase.Teacher,
	uu usecase.User,
	l *slog.Logger,
) {
	r := &teacherRoute{rmq, ut, uu, l}
	{
		router.protected.POST("/teacher/", r.createTeacher)
		router.public.GET("/teachers/", r.findAll)
	}
}

// SignUp

type createTeacherRequest dto.CreateTeacher

type createTeacherResponse *entity.Teacher

func (r *teacherRoute) createTeacher(c *gin.Context) {
	var body createTeacherRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	teacher, err := r.ut.Create(c.Request.Context(), dto.CreateTeacher(body))
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createTeacherResponse(teacher))

	response, err := json.Marshal(teacher)
	if err != nil {
		return
	}

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.teacher.created",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}

// FindAll

type findAllTeacherResponse struct {
	Teachers []entity.Teacher `json:"teachers"`
}

func (r *teacherRoute) findAll(c *gin.Context) {
	teachers, err := r.ut.FindAll(c.Request.Context())
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllTeacherResponse{Teachers: teachers})
}
