package database

import "net/url"

// ConnectionOptions manages the connection parameters needed for an SQL connection string
type ConnectionOptions struct {
	Name     string `json:"name"`
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Instance string `json:"instance"`
}

// ConnectionString creates an sql connection string
func (co *ConnectionOptions) ConnectionString() string {
	return (&url.URL{
		Scheme: co.Driver,
		User:   url.UserPassword(co.Username, co.Password),
		Host:   co.Host,
		Path:   co.Instance,
	}).String()
}
