package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/pkg/config"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
)

func main() {
	cfg := config.MustLoad()
	db, err := postgres.New(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	run := flag.String("run", "", "")

	flag.Parse()

	switch *run {
	case "create":
		create(db)
		break
	case "drop":
		drop(db)
		break
	case "reset":
		reset(db)
		break
	default:
		panic("?")
	}
}

func create(db *postgres.Postgres) error {
	if err := db.Exec(
		"CREATE TYPE user_role AS ENUM ('ADMIN', 'STUDENT', 'TEACHER', 'EMPLOYEE', 'MODERATOR')",
	).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(
		&entity.Faculty{},
		&entity.Department{},
		&entity.User{},
		&entity.Teacher{},
		&entity.Group{},
		&entity.Student{},
	); err != nil {
		return err
	}
	return nil
}

func drop(db *postgres.Postgres) error {
	tables := []string{"faculties", "departments", "users", "teachers", "groups", "students"}
	for _, t := range tables {
		if err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", t)).Error; err != nil {
			return err
		}
	}
	if err := db.Exec("DROP TYPE IF EXISTS user_role").Error; err != nil {
		return err
	}
	return nil
}

func reset(db *postgres.Postgres) {
	if err := drop(db); err != nil {
		panic(err)
	}

	if err := create(db); err != nil {
		panic(err)
	}
}
