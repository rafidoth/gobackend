package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/rafidoth/goback/internals/app"
)

func GetRouter(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.Checkhealth)
	return r
}
