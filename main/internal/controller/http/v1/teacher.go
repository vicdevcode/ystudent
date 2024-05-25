package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sethvargo/go-password/password"

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

type createTeacherResponse dto.TeacherResponse

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

	password, err := password.Generate(8, 8, 0, false, false)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(dto.TeacherUserResponse{
		User:     &teacher.User,
		Password: password,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createTeacherResponse{
		Teacher:  teacher,
		Password: password,
	})

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
	page, err := GetPage(c.Query("page"), c.Query("count"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	teachers, err := r.ut.FindAll(c.Request.Context(), page)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllTeacherResponse{Teachers: teachers})
}
