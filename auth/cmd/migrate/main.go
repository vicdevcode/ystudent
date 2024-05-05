package main

import (
	"log"

	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/config"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

func main() {
	cfg := config.MustLoad()
	db, err := postgres.New(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(
		&entity.Admin{},
		&entity.User{},
		&entity.Faculty{},
		&entity.Teacher{},
		&entity.Group{},
		&entity.Student{},
	)
}
