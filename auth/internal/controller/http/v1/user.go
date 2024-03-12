package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type userRoute struct {
	u usecase.User
	l *slog.Logger
}

func newUser(handler *gin.RouterGroup, u usecase.User, l *slog.Logger) {
	r := &userRoute{u, l}
	h := handler.Group("/user")
	{
		h.GET("/", r.findAll)
	}
}

// FindAll
type findAllUserResponse struct {
	Users []dto.UserResponse `json:"users"`
}

func (r *userRoute) findAll(c *gin.Context) {
	users, err := r.u.FindAll(c.Request.Context())
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
	c.JSON(http.StatusOK, findAllUserResponse{Users: userResponse})
}
