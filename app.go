package main

import (
	"net/http"
)

// NewApp allocates and return a new App
func NewApp() *App {
	env := NewEnv()
	return &App{
		router:  http.NewServeMux(),
		logging: env.Logging(),
	}
}

// App contains all the applications variables
type App struct {
	router  *http.ServeMux
	logging bool
}
