package middleware

import (
	"github.com/njwarm/celeritas"
	"github.com/njwarm/golaravel/data"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
