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

func NewStudent(handler *gin.RouterGroup, us usecase.Student, uu usecase.User, l *slog.Logger) {
	r := &studentRoute{us, uu, l}
	h := handler.Group("/student")
	{
		h.POST("/", r.signUp)
	}
}

// SignUp

type signUpStudentRequest struct {
	dto.CreateUserWithStudent
}

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
	var body signUpStudentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		errorJSONResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	student, err := r.us.SignUp(c.Request.Context(), body.CreateUserWithStudent)
	if err != nil {
		r.l.Error("sign up", slog.Any(err.Error(), err))
		errorResponse(c, http.StatusInternalServerError, "sign up")
		return
	}
	user, err := r.uu.FindOne(c.Request.Context(), student.UserID)
	response := signUpStudentResponse{User: user}
	response.Student = signUpGoodStudent{Student: student}

	c.JSON(http.StatusOK, response)
}
