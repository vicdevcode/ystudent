package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type studentRoute struct {
	us usecase.Student
	uu usecase.User
	l  *slog.Logger
}

func newStudent(handler *gin.RouterGroup, us usecase.Student, uu usecase.User, l *slog.Logger) {
	r := &studentRoute{us, uu, l}
	h := handler.Group("/student")
	{
		h.POST("/", r.signUp)
	}
}

// SignUp

type signUpStudentResponse struct {
	*entity.User
	Student signUpGoodStudent `json:"student"`
	dto.CUD
}

type signUpGoodStudent struct {
	*entity.Student
	GroupID interface{} `json:"group_id,omitempty"`
	UserID  interface{} `json:"user_id,omitempty"`
	dto.CUD
}

func (r *studentRoute) signUp(c *gin.Context) {
	var body dto.CreateUserAndStudent

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	student, err := r.us.SignUp(c.Request.Context(), body)
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}
	user, err := r.uu.FindOne(c.Request.Context(), entity.User{
		ID: student.UserID,
	})
	response := signUpStudentResponse{User: user}
	response.Student = signUpGoodStudent{Student: student}

	c.JSON(http.StatusOK, response)
}
