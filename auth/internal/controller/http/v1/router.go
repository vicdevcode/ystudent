package v1

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

func NewRouter(handler *gin.Engine, l *slog.Logger, uc usecase.UseCases) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/api/v1")
	{
		NewUser(h, uc.UserUseCase, l)
		NewStudent(h, uc.StudentUseCase, uc.UserUseCase, l)
	}
}
