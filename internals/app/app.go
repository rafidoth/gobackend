package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rafidoth/goback/internals/api"
	"github.com/rafidoth/goback/internals/store"
	"github.com/rafidoth/goback/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "logging : ", log.Ldate|log.Ltime)

	// grabbing data store from data layer
	workoutStore := store.NewPostgresWorkoutStore(pgDB)

	// injecting data from data layer to app layer
	workoutHandler := api.NewWorkoutHandler(workoutStore)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}
	return app, nil
}

func (a *Application) Checkhealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
