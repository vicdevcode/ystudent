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

type teacherRoute struct {
	ut usecase.Teacher
	uu usecase.User
	uh usecase.Hash
	l  *slog.Logger
}

func newTeacher(handler *gin.RouterGroup, ut usecase.Teacher, uu usecase.User, uh usecase.Hash, l *slog.Logger) {
	r := &teacherRoute{ut, uu, uh, l}
	h := handler.Group("/teacher")
	{
		h.POST("/create-with-user", r.createTeacherWithUser)
	}
}

// SignUp

type createTeacherWithUserResponse dto.UserResponse

func (r *teacherRoute) createTeacherWithUser(c *gin.Context) {
	var body dto.CreateUserAndTeacher

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
		internalServerError(c, err.Error())
		return
	}

	teacher, err := r.ut.Create(c.Request.Context(), dto.CreateTeacher{
		UserID: user.ID,
	})
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}
	user, err = r.uu.FindOne(c.Request.Context(), entity.User{
		ID: teacher.UserID,
	})

	c.JSON(http.StatusOK, createTeacherWithUserResponse{User: user})
}
