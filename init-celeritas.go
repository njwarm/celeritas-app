package main

import (
	"log"
	"os"

	"github.com/njwarm/celeritas"
	"github.com/njwarm/golaravel/data"
	"github.com/njwarm/golaravel/handlers"
	"github.com/njwarm/golaravel/middleware"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}

	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	myHandlers := &handlers.Handlers{
		App: cel,
	}

	app := &application{
		App:        cel,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)

	myHandlers.Models = app.Models

	app.Middleware.Models = app.Models

	return app
}
