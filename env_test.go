package main

import (
	"testing"
)

func lookupEnvMock(value string, ok bool) func(key string) (string, bool) {
	return func(key string) (string, bool) {
		return value, ok
	}
}

var env Env = Env{
	lookupEnv: lookupEnvMock("", false),
}

func TestLogging(t *testing.T) {
	// should return false if the env variable does not exists
	want := false
	got := env.Logging()
	if got != want {
		t.Errorf("env.Logging() = %v, want %v\n", got, want)
	}

	// should return false if the env variable exists and has value "false"
	env.lookupEnv = lookupEnvMock("false", true)
	want = false
	got = env.Logging()
	if got != want {
		t.Errorf("env.Logging() = %v, want %v\n", got, want)
	}

	// should return true if the variable existist and the value is not "false"
	env.lookupEnv = lookupEnvMock("randomValue", true)
	got = env.Logging()
	want = true
	if got != want {
		t.Errorf("env.Logging() = %v, want %v\n", got, want)
	}
}
