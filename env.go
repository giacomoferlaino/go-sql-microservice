package main

import (
	"os"
	"strings"
)

// NewEnv allocates and returns a new Env.
func NewEnv() *Env {
	return &Env{
		lookupEnv: os.LookupEnv,
	}
}

// Env reads environment variables values.
type Env struct {
	lookupEnv func(key string) (string, bool)
}

// Logging returns the values of the "logging" environment variables.
func (env *Env) Logging() bool {
	logging, ok := env.lookupEnv("logging")
	if ok == false {
		return false
	}
	if strings.ToLower(logging) == "false" {
		return false
	}
	return true
}
