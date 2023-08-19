package main

import (
	"github.com/demig00d/zakaty-service/config"
	"github.com/demig00d/zakaty-service/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
