package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type userRoute struct {
	u usecase.User
	l *slog.Logger
}

func NewUser(handler *gin.RouterGroup, u usecase.User, l *slog.Logger) {
	r := &userRoute{u, l}
	h := handler.Group("/user")
	{
		h.GET("/", r.findAll)
		h.POST("/", r.signUp)
	}
}

// FindAll
type findAllUserResponse struct {
	Users []entity.User `json:"users"`
}

func (r *userRoute) findAll(c *gin.Context) {
	users, err := r.u.FindAll(c.Request.Context())
	if err != nil {
		r.l.Error("http-v1-user-find-all", slog.Any("error", err))
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, findAllUserResponse{users})
}

// SignUp

type signUpUserRequest struct {
	dto.CreateUser
}

type signUpUserResponse struct {
	dto.User
}

func (r *userRoute) signUp(c *gin.Context) {
	var body signUpUserRequest

	if err := c.BindJSON(&body); err != nil {
		errorResponse(c, http.StatusBadRequest, "sign up")
		return
	}

	user, err := r.u.SignUp(c.Request.Context(), body.CreateUser)
	if err != nil {
		r.l.Error("sign up", slog.Any(err.Error(), err))
		errorResponse(c, http.StatusInternalServerError, "sign up")
		return
	}
	c.JSON(http.StatusOK, user)
}
