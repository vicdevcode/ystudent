package main

import (
	"github.com/vicdevcode/ystudent/auth/internal/app"
	"github.com/vicdevcode/ystudent/auth/pkg/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
