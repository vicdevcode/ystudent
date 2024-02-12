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

func NewTeacher(handler *gin.RouterGroup, ut usecase.Teacher, uu usecase.User, l *slog.Logger) {
	r := &teacherRoute{ut, uu, l}
	h := handler.Group("/teacher")
	{
		h.POST("/", r.signUp)
	}
}

// SignUp

type signUpTeacherRequest struct {
	dto.CreateUser
}

type signUpTeacherResponse struct {
	*entity.User
	Teacher signUpGoodTeacher `json:"teacher"`
	dto.CUD
}

type signUpGoodTeacher struct {
	*entity.Teacher
	UserID interface{} `json:"user_id,omitempty"`
	dto.CUD
}

func (r *teacherRoute) signUp(c *gin.Context) {
	var body signUpTeacherRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		errorJSONResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	teacher, err := r.ut.SignUp(c.Request.Context(), body.CreateUser)
	if err != nil {
		r.l.Error("sign up", slog.Any(err.Error(), err))
		errorResponse(c, http.StatusInternalServerError, "sign up")
		return
	}
	user, err := r.uu.FindOne(c.Request.Context(), teacher.UserID)
	response := signUpTeacherResponse{User: user}
	response.Teacher = signUpGoodTeacher{Teacher: teacher}

	c.JSON(http.StatusOK, response)
}
