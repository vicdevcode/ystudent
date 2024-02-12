package main

import (
	"log"

	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/config"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

func main() {
	cfg := config.MustLoad()
	db, err := postgres.New((*postgres.Config)(&cfg.DB))
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(
		&entity.Faculty{},
		&entity.User{},
		&entity.Teacher{},
		&entity.Group{},
		&entity.Student{},
	)
}
