package main

import (
	"log"

	"github.com/millirud/easy-content/video-api/config"
	"github.com/millirud/easy-content/video-api/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
