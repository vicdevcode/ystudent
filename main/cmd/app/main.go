package main

import (
	"github.com/vicdevcode/ystudent/main/internal/app"
	"github.com/vicdevcode/ystudent/main/pkg/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
