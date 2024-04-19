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

type studentRoute struct {
	us  usecase.Student
	uu  usecase.User
	uh  usecase.Hash
	rmq *RabbitMQ
	l   *slog.Logger
}

func newStudent(
	handler *gin.RouterGroup,
	us usecase.Student,
	uu usecase.User,
	uh usecase.Hash,
	rmq *RabbitMQ,
	l *slog.Logger,
) {
	r := &studentRoute{us, uu, uh, rmq, l}
	h := handler.Group("/student")
	{
		h.POST("/create-with-user", r.createStudentWithUser)
		h.GET("/", r.findAll)
	}
}

// CreateWithUser

type createStudentWithUserResponse dto.UserResponse

func (r *studentRoute) createStudentWithUser(c *gin.Context) {
	var body dto.CreateUserAndStudent

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	// password, err := password.Generate(8, 8, 0, false, false)
	// if err != nil {
	// 	internalServerError(c, err.Error())
	// 	return
	// }

	hashedPassword, err := r.uh.HashPassword("123123123")
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	createUser := dto.CreateUser{
		Fio:      dto.Fio(body.CreateUserWithoutPassword.Fio),
		Email:    body.CreateUserWithoutPassword.Email,
		Password: hashedPassword,
	}

	user, err := r.uu.Create(c.Request.Context(), createUser)
	if err != nil {
		r.l.Error(err.Error())
		badRequest(c, err.Error())
		return
	}
	_, err = r.us.Create(c.Request.Context(), dto.CreateStudent{
		UserID:  user.ID,
		GroupID: body.GroupID,
		Leader:  body.Leader,
	})
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	user, err = r.uu.FindOne(c.Request.Context(), entity.User{ID: user.ID})

	c.JSON(http.StatusOK, createStudentWithUserResponse{User: user})

	response, err := json.Marshal(user)
	if err != nil {
		return
	}

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"auth.student.created",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}

// Create

type createStudentResponse dto.StudentResponse

func (r *studentRoute) createStudent(c *gin.Context) {
	var body dto.CreateStudent

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	student, err := r.us.Create(c.Request.Context(), dto.CreateStudent{
		UserID:  body.UserID,
		GroupID: body.GroupID,
		Leader:  body.Leader,
	})
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
		"auth.student.created",
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
	Users []dto.UserResponse `json:"users"`
}

func (r *studentRoute) findAll(c *gin.Context) {
	users, err := r.us.FindAll(c.Request.Context())
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	userResponse := make([]dto.UserResponse, len(users), len(users))
	for i, user := range users {
		currentUser := user
		userResponse[i] = dto.UserResponse{User: &currentUser}
	}
	c.JSON(http.StatusOK, findAllStudentUserResponse{Users: userResponse})
}
