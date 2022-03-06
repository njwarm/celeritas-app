package main

import (
	"github.com/njwarm/celeritas"
	"github.com/njwarm/golaravel/data"
	"github.com/njwarm/golaravel/handlers"
	"github.com/njwarm/golaravel/middleware"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
