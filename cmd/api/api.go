// Copyright (c) 2025. guimochila.com. Continuous Learning.

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/guimochila/social/internal/config"
	"github.com/guimochila/social/internal/store"
)

const (
	MAX_WRITE_TIMEOUT = 30
	MAX_READ_TIMEOUT  = 10
)

type application struct {
	config config.Config
	store  store.Storage
	db     config.DBConfig
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second)) //nolint: mnd

	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthz", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.GetAddress(),
		Handler:      mux,
		WriteTimeout: time.Second * MAX_WRITE_TIMEOUT,
		ReadTimeout:  time.Second * MAX_READ_TIMEOUT,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at %s", app.config.GetAddress())

	return srv.ListenAndServe()
}
