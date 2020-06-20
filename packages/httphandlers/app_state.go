package httphandlers

import "github.com/giacomoferlaino/go-sql-microservice/packages/database"

// AppState represents the current application state
type AppState interface {
	Databases() *database.SyncMap
	Logging() bool
}
