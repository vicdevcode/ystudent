package app

import (
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
)

func migrate(db *postgres.Postgres) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Faculty{},
		&entity.Teacher{},
		&entity.Group{},
		&entity.Student{},
	); err != nil {
		return err
	}
	return nil
}
