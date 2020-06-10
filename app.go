package sqlmicroservice

import (
	"net/http"
)

// NewApp allocates and return a new App
func NewApp() *App {
	env := NewEnv()
	return &App{
		Router:  http.NewServeMux(),
		Logging: env.Logging(),
	}
}

// App contains all the applications variables
type App struct {
	Router  *http.ServeMux
	Logging bool
}
