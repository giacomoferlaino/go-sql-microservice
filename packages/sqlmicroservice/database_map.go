package sqlmicroservice

import (
	"database/sql"
	"sync"
)

type mutexMap interface {
	Load(key interface{}) (interface{}, bool)
	Store(key interface{}, value interface{})
	Delete(key interface{})
	Range(func(key interface{}, value interface{}) bool)
}

// NewDatabaseMap allocates and return a new database map.
func NewDatabaseMap() *DatabaseMap {
	return &DatabaseMap{
		connections: &sync.Map{},
	}
}

// DatabaseMap manages a map of database connections protected by a mutex.
type DatabaseMap struct {
	connections mutexMap
}

// Load return an existing connection given its key if it exists
func (db *DatabaseMap) Load(key string) (*sql.DB, bool) {
	connection, ok := db.connections.Load(key)
	if connection == nil {
		return nil, ok
	}
	return connection.(*sql.DB), ok
}

// Store saves a database connection on a specific key
func (db *DatabaseMap) Store(key string, connection *sql.DB) {
	db.connections.Store(key, connection)
}

// Delete removes an existing connection given its key
func (db *DatabaseMap) Delete(key string) {
	db.connections.Delete(key)
}

// Range is used to loop over all the database connections
// and their relative keys
func (db *DatabaseMap) Range(f func(key string, connection *sql.DB) bool) {
	db.connections.Range(func(key interface{}, value interface{}) bool {
		return f(key.(string), value.(*sql.DB))
	})
}
