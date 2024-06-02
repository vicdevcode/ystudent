package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"

	"github.com/vicdevcode/ystudent/main/internal/controller/amqp/v1/consumer"
	v1 "github.com/vicdevcode/ystudent/main/internal/controller/http/v1"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/internal/usecase"
	"github.com/vicdevcode/ystudent/main/pkg/config"
	"github.com/vicdevcode/ystudent/main/pkg/httpserver"
	"github.com/vicdevcode/ystudent/main/pkg/logger"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
	"github.com/vicdevcode/ystudent/main/pkg/rabbitmq"
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

	// Set Admin
	var admin *entity.User
	if err := db.Where("email = ?", cfg.Admin.Email).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			admin := &entity.User{
				Email:    cfg.Admin.Email,
				RoleType: entity.ADMIN,
			}
			if err = db.Create(admin).Error; err != nil {
				log.Error(err.Error())
				return
			}

			adminJSON, err := json.Marshal(admin)
			if err != nil {
				log.Error("can not marshal admin")
				return
			}

			ch.PublishWithContext(
				context.Background(),
				cfg.RabbitMQ.ExchangeName,
				"main.admin.created",
				false,
				false,
				amqp091.Publishing{
					ContentType: "application/json",
					Body:        adminJSON,
				},
			)
		} else {
			log.Error(err.Error())
			return
		}
	}
	log.Info("Admin available")

	// HTTP SERVER
	gin.SetMode(gin.ReleaseMode)
	if cfg.Env == "local" {
		gin.SetMode(gin.DebugMode)
	}
	handler := gin.New()
	v1.NewRouter(handler, ch, cfg.RabbitMQ.ExchangeName, cfg.Env, log, usecases)
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
