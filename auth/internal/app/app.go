package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/controller/amqp/v1/consumer"
	v1 "github.com/vicdevcode/ystudent/auth/internal/controller/http/v1"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
	"github.com/vicdevcode/ystudent/auth/pkg/config"
	"github.com/vicdevcode/ystudent/auth/pkg/httpserver"
	"github.com/vicdevcode/ystudent/auth/pkg/logger"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
	"github.com/vicdevcode/ystudent/auth/pkg/rabbitmq"
)

func Run(cfg *config.Config) {
	log := logger.New(cfg.Env)

	log.Info(fmt.Sprintf("Starting server at Port: %s", cfg.HTTP.Port))

	db, err := postgres.New(&cfg.DB)
	if err != nil {
		logger.Fatal(log, "Failed connect to sqlite:", err)
	}
	log.Info("Connected to postgres")

	// AutoMigrate
	if cfg.Env != "local" {
		if err := migrate(db); err != nil {
			logger.Fatal(log, "Failed to migrate:", err)
		}
		log.Info("Migration completed successfully")
	}

	// UseCases
	usecases := usecase.New(cfg, db)

	// RabbitMQ
	conn, ch, _, delivery := rabbitmq.New(&cfg.RabbitMQ)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	consumer := consumer.New(
		conn,
		ch,
		ctx,
		delivery,
	)

	go consumer.Start(usecases, log)

	// HTTP SERVER
	gin.SetMode(gin.ReleaseMode)
	if cfg.Env == "local" {
		gin.SetMode(gin.DebugMode)
	}
	handler := gin.New()
	v1.NewRouter(handler, ch, cfg.RabbitMQ.ExchangeName, log, usecases)
	httpServer := httpserver.New(handler, (&cfg.HTTP))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Fatal(log, "app - Run - httpServer.Notify", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Fatal(log, "app - Run - httpServer.Shutdown", err)
	}
}
