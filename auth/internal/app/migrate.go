package app

import (
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

func migrate(db *postgres.Postgres) error {
	if err := db.AutoMigrate(
		&entity.Admin{},
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
