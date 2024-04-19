package app

import (
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/sqlite"
)

func migrate(db *sqlite.SQLite) error {
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
