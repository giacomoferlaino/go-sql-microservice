package database

import "testing"

func TestConnectionString(t *testing.T) {
	connectionOptions := ConnectionOptions{
		Driver:   "sqlserver",
		Username: "sa",
		Password: "password",
		Host:     "localhost",
		Instance: "master",
	}
	want := "sqlserver://sa:password@localhost/master"
	got := connectionOptions.ConnectionString()
	if got != want {
		t.Errorf("connectionOptions.ConnectionString() = %v, want %v\n", got, want)
	}
}
