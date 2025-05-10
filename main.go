package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/rafidoth/goback/internals/app"
	"github.com/rafidoth/goback/internals/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Backend Port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	app.Logger.Printf("app is running at port %d\n", port)

	r := routes.GetRouter(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  0,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}

}
