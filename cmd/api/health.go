// Copyright (c) 2025. guimochila.com. Continuous Learning.

package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
