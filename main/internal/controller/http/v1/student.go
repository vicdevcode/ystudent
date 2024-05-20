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

type studentRoute struct {
	rmq *RabbitMQ
	us  usecase.Student
	uu  usecase.User
	l   *slog.Logger
}

func newStudent(
	router *router,
	rmq *RabbitMQ,
	us usecase.Student,
	uu usecase.User,
	l *slog.Logger,
) {
	r := &studentRoute{rmq, us, uu, l}
	{
		router.protected.POST("/student/", r.createStudent)
		router.public.GET("/students/", r.findAll)
	}
}

// Create
type createStudentRequest dto.CreateStudent

type createStudentResponse dto.StudentResponse

func (r *studentRoute) createStudent(c *gin.Context) {
	var body dto.CreateStudent

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	student, err := r.us.Create(c.Request.Context(), dto.CreateStudent(body))
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createStudentResponse{
		Student: student,
	})

	response, err := json.Marshal(student)
	if err != nil {
		return
	}

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.student.created",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}

// FindAll

type findAllStudentUserResponse struct {
	Students []entity.Student `json:"students"`
}

func (r *studentRoute) findAll(c *gin.Context) {
	students, err := r.us.FindAll(c.Request.Context())
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllStudentUserResponse{Students: students})
}
