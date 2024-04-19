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

type facultyRoute struct {
	u   usecase.Faculty
	l   *slog.Logger
	rmq *RabbitMQ
}

func newFaculty(
	public *gin.RouterGroup,
	protected *gin.RouterGroup,
	u usecase.Faculty, rmq *RabbitMQ, l *slog.Logger,
) {
	r := &facultyRoute{u, l, rmq}
	{
		protected.POST("/faculty/", r.create)
		public.GET("/faculty/", r.findAll)
	}
}

type createFacultyResponse struct {
	*entity.Faculty
	dto.CUD
}

func (r *facultyRoute) create(c *gin.Context) {
	var body dto.CreateFaculty

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	faculty, err := r.u.Create(c.Request.Context(), body)
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(faculty)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, createFacultyResponse{Faculty: faculty})

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"auth.faculty.created",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}

type findAllFacultyResponse struct {
	Faculties []entity.Faculty `json:"faculties"`
}

func (r *facultyRoute) findAll(c *gin.Context) {
	faculties, err := r.u.FindAll(c.Request.Context())
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllFacultyResponse{Faculties: faculties})
}
