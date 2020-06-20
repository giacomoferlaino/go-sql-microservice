package database

import (
	"database/sql"
	"encoding/json"
	"sync"
)

// MutexMap represents a map safe for concurrency
type MutexMap interface {
	Load(key interface{}) (interface{}, bool)
	Store(key interface{}, value interface{})
	Delete(key interface{})
	Range(func(key interface{}, value interface{}) bool)
}

// NewSyncMap allocates and return a new database map.
func NewSyncMap() *SyncMap {
	return &SyncMap{
		connections: &sync.Map{},
	}
}

// SyncMap manages a map of database connections protected by a mutex.
type SyncMap struct {
	connections MutexMap
}

// Load return an existing connection given its key if it exists
func (db *SyncMap) Load(key string) (*sql.DB, bool) {
	connection, ok := db.connections.Load(key)
	if connection == nil {
		return nil, ok
	}
	return connection.(*sql.DB), ok
}

// Store saves a database connection on a specific key
func (db *SyncMap) Store(key string, connection *sql.DB) {
	db.connections.Store(key, connection)
}

// Delete removes an existing connection given its key
func (db *SyncMap) Delete(key string) {
	db.connections.Delete(key)
}

// Range is used to loop over all the database connections
// and their relative keys
func (db *SyncMap) Range(f func(key string, connection *sql.DB) bool) {
	db.connections.Range(func(key interface{}, value interface{}) bool {
		return f(key.(string), value.(*sql.DB))
	})
}

// MarshalJSON converts data into a valid JSON format
func (db *SyncMap) MarshalJSON() ([]byte, error) {
	tmpMap := map[string]sql.DBStats{}
	db.connections.Range(func(key, value interface{}) bool {
		tmpMap[key.(string)] = value.(*sql.DB).Stats()
		return true
	})
	return json.Marshal(tmpMap)
}
