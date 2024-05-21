package app

import (
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

func migrate(db *postgres.Postgres) error {
	if err := db.AutoMigrate(
		&entity.User{},
	); err != nil {
		return err
	}
	return nil
}
