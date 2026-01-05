package main

import (
	"log"

	"memo-go/services/pos/internal/app"
	"memo-go/services/pos/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatalf("run app: %v", err)
	}
}
