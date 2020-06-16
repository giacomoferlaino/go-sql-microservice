package sqlmicroservice

import (
	"os"
	"strings"
)

// NewEnv allocates and returns a new Env.
func NewEnv() *Env {
	return &Env{
		LookupEnv: os.LookupEnv,
	}
}

// Env reads environment variables values.
type Env struct {
	LookupEnv func(key string) (string, bool)
}

// Logging returns the values of the "logging" environment variables.
func (env *Env) Logging() bool {
	logging, ok := env.LookupEnv("logging")
	if ok == false {
		return false
	}
	if strings.ToLower(logging) == "false" {
		return false
	}
	return true
}
