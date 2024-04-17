package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Applicaiton) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/quiz", app.Quiz)

	return r
}
