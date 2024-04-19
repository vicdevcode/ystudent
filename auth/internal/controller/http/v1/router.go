package v1

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type StartMessage struct {
	Message string `json:"message"`
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
	handler.Use(cors.Default())

	public := handler.Group("/api/v1")
	private := handler.Group("/api/v1")
	protected := handler.Group("/api/v1")
	private.Use(jwtCheckMiddleware(uc.JwtUseCase))
	protected.Use(adminCheckMiddleware(uc.JwtUseCase, uc.AdminUseCase))

	message, err := json.Marshal(StartMessage{Message: "microservice is started"})
	if err != nil {
		l.Error(err.Error())
		return
	}

	rmq := newRabbitMQ(ch, exchange)
	rmq.ch.PublishWithContext(
		context.Background(),
		exchange,
		"auth.start",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)

	{
		newUser(protected, uc.UserUseCase, l)
		newStudent(protected, uc.StudentUseCase, uc.UserUseCase, uc.HashUseCase, rmq, l)
		newTeacher(protected, uc.TeacherUseCase, uc.UserUseCase, uc.HashUseCase, rmq, l)
		newFaculty(public, protected, uc.FacultyUseCase, rmq, l)
		newGroup(public, protected, uc.GroupUseCase, rmq, l)
		newAuth(public, private, uc.AdminUseCase, uc.UserUseCase, uc.HashUseCase, uc.JwtUseCase, l)
	}
}
