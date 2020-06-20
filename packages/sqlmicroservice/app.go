// Package sqlmicroservice provides the global functionalities needed to run the microservice.
package sqlmicroservice

import (
	"net/http"

	"github.com/giacomoferlaino/go-sql-microservice/packages/database"
	"github.com/giacomoferlaino/go-sql-microservice/packages/httphandlers"
)

// NewApp allocates and return a new App.
func NewApp() *App {
	env := NewEnv()
	return &App{
		databases: database.NewSyncMap(),
		router:    http.NewServeMux(),
		logging:   env.Logging(),
	}
}

// App contains all the applications variables.
type App struct {
	databases *database.SyncMap
	router    *http.ServeMux
	logging   bool
}

// Router gets the router property
func (app *App) Router() *http.ServeMux {
	return app.router
}

// Databases gets the databases property
func (app *App) Databases() *database.SyncMap {
	return app.databases
}

// Logging gets the logging property
func (app *App) Logging() bool {
	return app.logging
}

// DefaultHandlers loads the default http handlers.
func (app *App) DefaultHandlers() {
	app.router.Handle("/connection", httphandlers.NewConnectionHandler(app))
}
