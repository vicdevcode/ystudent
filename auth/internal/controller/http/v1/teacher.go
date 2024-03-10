package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type teacherRoute struct {
	ut usecase.Teacher
	uu usecase.User
	l  *slog.Logger
}

func newTeacher(handler *gin.RouterGroup, ut usecase.Teacher, uu usecase.User, l *slog.Logger) {
	r := &teacherRoute{ut, uu, l}
	h := handler.Group("/teacher")
	{
		h.POST("/", r.signUp)
	}
}

// SignUp

type signUpTeacherResponse struct {
	*entity.User
	Teacher struct {
		*entity.Teacher
		UserID interface{} `json:"user_id,omitempty"`
		dto.CUD
	} `json:"teacher"`
	dto.CUD
}

func (r *teacherRoute) signUp(c *gin.Context) {
	var body dto.CreateUserAndTeacher

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	teacher, err := r.ut.SignUp(c.Request.Context(), body.CreateUser)
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}
	user, err := r.uu.FindOne(c.Request.Context(), entity.User{
		ID: teacher.UserID,
	})
	response := signUpTeacherResponse{User: user}
	response.Teacher.Teacher = teacher

	c.JSON(http.StatusOK, response)
}
