package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	*gorm.DB
}

type Config struct {
	Path string `yaml:"path" env-required:"true"`
}

func New(cfg *Config) (*SQLite, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &SQLite{db}, nil
}
