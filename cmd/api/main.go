// Copyright (c) 2025. guimochila.com. Continuous Learning.

package main

import (
	"log"

	"github.com/guimochila/social/internal/config"
	"github.com/guimochila/social/internal/db"
	"github.com/guimochila/social/internal/store"
)

func main() {
	var cfg config.Config
	if err := config.FromEnv(&cfg); err != nil {
		log.Fatal("error loading configuration")
	}

	db, err := db.New(
		cfg.DB.DSN(),
		cfg.DB.GetMaxIdleTime(),
		cfg.DB.MaxOpenConns,
		cfg.DB.MaxIdleConns,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Fatalf("Application failed: %v\n", err)
	}
}
