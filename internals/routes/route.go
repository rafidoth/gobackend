package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/rafidoth/goback/internals/app"
)

func GetRouter(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.Checkhealth)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)

	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkout)

	r.Delete("/workouts/{id}", app.WorkoutHandler.HandleDeleteWOrkout)
	return r
}
