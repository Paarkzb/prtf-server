package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *applicaiton) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", app.quiz)

	return r
}
