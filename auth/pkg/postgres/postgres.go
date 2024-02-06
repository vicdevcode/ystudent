package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	*gorm.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func New(cfg *Config) (*Postgres, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}
