package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sethvargo/go-password/password"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/internal/usecase"
)

type studentRoute struct {
	rmq *RabbitMQ
	us  usecase.Student
	uu  usecase.User
	ug  usecase.Group
	l   *slog.Logger
}

func newStudent(
	router *router,
	rmq *RabbitMQ,
	us usecase.Student,
	uu usecase.User,
	ug usecase.Group,
	l *slog.Logger,
) {
	r := &studentRoute{rmq, us, uu, ug, l}
	{
		router.protected.POST("/student/", r.create)
		router.protected.PUT("/student/:id", r.update)
		router.protected.DELETE("/student/:id", r.delete)
		router.public.GET("/students/", r.findAll)
	}
}

// Create
type createStudentRequest dto.CreateStudent

type createStudentResponse dto.StudentResponse

func (r *studentRoute) create(c *gin.Context) {
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

	password, err := password.Generate(8, 8, 0, false, false)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(dto.StudentUserResponse{
		GroupID:  student.GroupID,
		User:     &student.User,
		Password: password,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createStudentResponse{
		Student:  student,
		Password: password,
	})

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

type findAllStudent struct {
	*entity.Student
	GroupName string `json:"group_name"`
}

type findAllStudentUserResponse struct {
	Students []findAllStudent `json:"students"`
	dto.Page
}

func (r *studentRoute) findAll(c *gin.Context) {
	page, err := GetPage(c.Query("page"), c.Query("count"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	students, err := r.us.FindAll(c.Request.Context(), page)
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	var response []findAllStudent
	for _, student := range students {
		group, err := r.ug.FindOne(c.Request.Context(), entity.Group{
			ID: student.GroupID,
		})
		if err != nil {
			r.l.Error(err.Error())
			internalServerError(c, err.Error())
			return
		}
		response = append(response, findAllStudent{
			Student:   &student,
			GroupName: group.Name,
		})
	}

	c.JSON(http.StatusOK, findAllStudentUserResponse{Students: response, Page: page})
}

type updateStudentRequest dto.UpdateStudentBody

type updateStudentResponse *entity.Student

func (r *studentRoute) update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	var body updateStudentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	student, err := r.us.FindOne(c.Request.Context(), entity.Student{
		ID: id,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	data := dto.UpdateUser{
		ID: student.User.ID,
	}
	if body.Firstname != "" {
		data.Firstname = body.Firstname
	}
	if body.Surname != "" {
		data.Surname = body.Surname
	}
	if body.Middlename != "" {
		data.Middlename = body.Middlename
	}
	if body.Email != "" {
		data.Email = body.Email
	}

	_, err = r.uu.Update(c.Request.Context(), data)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	if body.GroupID != uuid.Nil {
		r.l.Info("", slog.Any("", body))
		_, err := r.us.Update(c.Request.Context(), dto.UpdateStudent{
			ID:      student.ID,
			GroupID: body.GroupID,
		})
		if err != nil {
			internalServerError(c, err.Error())
			return
		}
	}

	student, err = r.us.FindOne(c.Request.Context(), entity.Student{
		ID: id,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateStudentResponse(student))
}

func (r *studentRoute) delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	if err := r.us.Delete(c.Request.Context(), id); err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, deleteGroupResponse{
		Message: "student was deleted",
	})
}
