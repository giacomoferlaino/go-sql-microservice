package database

import "net/url"

// ConnectionOptions manages the connection parameters needed for an SQL connection string
type ConnectionOptions struct {
	driver   string
	username string
	password string
	host     string
	instance string
}

func (co *ConnectionOptions) connectionString() string {
	return (&url.URL{
		Scheme: co.driver,
		User:   url.UserPassword(co.username, co.password),
		Host:   co.host,
		Path:   co.instance,
	}).String()
}
