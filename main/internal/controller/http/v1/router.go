package v1

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/vicdevcode/ystudent/main/internal/usecase"
)

type StartMessage struct {
	Message string `json:"message"`
}

type router struct {
	public    *gin.RouterGroup
	private   *gin.RouterGroup
	protected *gin.RouterGroup
	moderator *gin.RouterGroup
}

func NewRouter(
	handler *gin.Engine,
	ch *amqp.Channel,
	exchange string,
	l *slog.Logger,
	uc usecase.UseCases,
) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	public := handler.Group("/api/v1")
	private := handler.Group("/api/v1")
	moderator := handler.Group("/api/v1")
	protected := handler.Group("/api/v1")
	private.Use(userAuthCheckMiddleware(uc.JwtUseCase))
	protected.Use(adminAuthCheckMiddleware(uc.JwtUseCase))
	moderator.Use(moderatorAuthCheckMiddleware(uc.JwtUseCase))

	message, err := json.Marshal(StartMessage{Message: "microservice is started"})
	if err != nil {
		l.Error(err.Error())
		return
	}

	rmq := newRabbitMQ(ch, exchange)

	router := &router{
		public:    public,
		private:   private,
		protected: protected,
		moderator: moderator,
	}

	rmq.ch.PublishWithContext(
		context.Background(),
		rmq.exchange,
		"main.start",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)

	{
		newDepartment(router, rmq, uc.DepartmentUseCase, uc.FacultyUseCase, l)
		newModerator(router, rmq, uc.UserUseCase, l)
		newEmployee(router, rmq, uc.EmployeeUseCase, l)
		newStudent(router, rmq, uc.StudentUseCase, uc.UserUseCase, uc.GroupUseCase, l)
		newTeacher(router, rmq, uc.TeacherUseCase, uc.UserUseCase, l)
		newFaculty(router, rmq, uc.FacultyUseCase, l)
		newGroup(router, rmq, uc.GroupUseCase, uc.DepartmentUseCase, uc.TeacherUseCase, l)
	}
}
