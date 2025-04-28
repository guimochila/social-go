// Copyright (c) 2025. guimochila.com. Continuous Learning.

package main

import (
	"log"

	"github.com/guimochila/social/internal/config"
)

func main() {
	var cfg config.Config
	if err := config.FromEnv(&cfg); err != nil {
		log.Fatal("error loading configuration")
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Fatalf("Application failed: %v\n", err)
	}
}
