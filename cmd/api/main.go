// Copyright (c) 2025. guimochila.com. Continuous Learning.
package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Fatalf("Application failed: %v\n", err)
	}
}
