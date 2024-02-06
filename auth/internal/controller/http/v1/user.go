package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

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
