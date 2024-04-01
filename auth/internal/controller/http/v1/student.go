package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type studentRoute struct {
	us usecase.Student
	uu usecase.User
	uh usecase.Hash
	l  *slog.Logger
}

func newStudent(
	handler *gin.RouterGroup,
	us usecase.Student,
	uu usecase.User,
	uh usecase.Hash,
	l *slog.Logger,
) {
	r := &studentRoute{us, uu, uh, l}
	h := handler.Group("/student")
	{
		h.POST("/create-with-user", r.createStudentWithUser)
		h.GET("/", r.findAll)
	}
}

// CreateWithUser

type createStudentWithUserResponse struct {
	dto.UserResponse
	Password string `json:"password"`
}

func (r *studentRoute) createStudentWithUser(c *gin.Context) {
	var body dto.CreateUserAndStudent

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	password, err := password.Generate(8, 8, 0, false, false)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	hashedPassword, err := r.uh.HashPassword(password)
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

	c.JSON(http.StatusOK, createStudentWithUserResponse{
		UserResponse: dto.UserResponse{User: user},
		Password:     password,
	})
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
