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
		newUser(h, uc.UserUseCase, l)
		newStudent(h, uc.StudentUseCase, uc.UserUseCase, l)
		newTeacher(h, uc.TeacherUseCase, uc.UserUseCase, l)
		newFaculty(h, uc.FacultyUseCase, l)
		newGroup(h, uc.GroupUseCase, l)
		newAuth(h, uc.AdminUseCase, uc.HashUseCase, uc.JwtUseCase, l)
	}
}
