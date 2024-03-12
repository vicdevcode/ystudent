package v1

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

func NewRouter(handler *gin.Engine, l *slog.Logger, uc usecase.UseCases) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(cors.Default())

	h := handler.Group("/api/v1")
	private := handler.Group("/api/v1")
	protected := handler.Group("/api/v1")
	private.Use(jwtCheckMiddleware(uc.JwtUseCase))
	protected.Use(adminCheckMiddleware(uc.JwtUseCase, uc.AdminUseCase))
	{
		newUser(protected, uc.UserUseCase, l)
		newStudent(protected, uc.StudentUseCase, uc.UserUseCase, uc.HashUseCase, l)
		newTeacher(protected, uc.TeacherUseCase, uc.UserUseCase, uc.HashUseCase, l)
		newFaculty(h, uc.FacultyUseCase, l)
		newGroup(h, uc.GroupUseCase, l)
		newAuth(h, uc.AdminUseCase, uc.UserUseCase, uc.HashUseCase, uc.JwtUseCase, l)
	}
}
