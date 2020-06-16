package sqlmicroservice

import (
	"net/http"

	"github.com/giacomoferlaino/go-sql-microservice/packages/sqlmicroservice/httphandlers"
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

// DefaultHandlers loads the default http handlers
func (app App) DefaultHandlers() {
	app.Router.HandleFunc("/connect", httphandlers.ConnectHandler)
}
