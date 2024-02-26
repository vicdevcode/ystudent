package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"

	"github.com/vicdevcode/ystudent/auth/pkg/httpserver"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

type Config struct {
	Env            string            `yaml:"env"             env-required:"true"`
	ContextTimeout time.Duration     `yaml:"context_timeout" env-required:"true"`
	HTTP           httpserver.Config `yaml:"http"            env-required:"true"`
	DB             postgres.Config   `yaml:"postgres"        env-required:"true"`
	Admin          Admin             `yaml:"admin"           env-required:"true"`
	JWT            JWT               `yaml:"jwt"             env-required:"true"`
}

type Admin struct {
	Login    string `yaml:"login"    env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

type JWT struct {
	Secret           string        `yaml:"secret"             env-required:"true"`
	AccessExpiresAt  time.Duration `yaml:"access_expires_at"  env-required:"true"`
	RefreshExpiresAt time.Duration `yaml:"refresh_expires_at" env-required:"true"`
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
