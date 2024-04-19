package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"

	"github.com/vicdevcode/ystudent/auth/pkg/httpserver"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
	"github.com/vicdevcode/ystudent/auth/pkg/rabbitmq"
	"github.com/vicdevcode/ystudent/auth/pkg/sqlite"
)

type Config struct {
	Env            string            `yaml:"env"             env-required:"true"`
	ContextTimeout time.Duration     `yaml:"context_timeout" env-required:"true"`
	FillDatabase   *bool             `yaml:"fill_database"`
	HTTP           httpserver.Config `yaml:"http"            env-required:"true"`
	DB             postgres.Config   `yaml:"postgres"`
	Admin          Admin             `yaml:"admin"           env-required:"true"`
	JWT            JWT               `yaml:"jwt"             env-required:"true"`
	RabbitMQ       rabbitmq.Config   `yaml:"rabbitmq"        env-required:"true"`
	SQLite         sqlite.Config     `yaml:"sqlite"`
}

type Admin struct {
	Login    string `yaml:"login"    env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

type JWT struct {
	AccessTokenSecret  string        `yaml:"access_token_secret"  env-required:"true"`
	RefreshTokenSecret string        `yaml:"refresh_token_secret" env-required:"true"`
	AccessExpiresAt    time.Duration `yaml:"access_expires_at"    env-required:"true"`
	RefreshExpiresAt   time.Duration `yaml:"refresh_expires_at"   env-required:"true"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
