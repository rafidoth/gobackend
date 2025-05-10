package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "logging : ", log.Ldate|log.Ltime)

	app := &Application{
		Logger: logger,
	}
	return app, nil
}

func (a *Application) Checkhealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available")
}
