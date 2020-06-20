package database

import "testing"

func TestConnectionString(t *testing.T) {
	connectionOptions := ConnectionOptions{
		driver:   "sqlserver",
		username: "sa",
		password: "password",
		host:     "localhost",
		instance: "master",
	}
	want := "sqlserver://sa:password@localhost/master"
	got := connectionOptions.connectionString()
	if got != want {
		t.Errorf("connectionOptions.connectionString() = %v, want %v\n", got, want)
	}
}
