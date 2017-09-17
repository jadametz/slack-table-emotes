package main

import (
	"os"
	"testing"
)

func TestGetEnvExists(t *testing.T) {
	expected := "foo"
	fallback := "bar"
	os.Setenv("TESTENVKEY", expected)
	envValue := getEnv("TESTENVKEY", fallback)
	if envValue != expected {
		t.Errorf("Return value was incorrect, got: %s, want: %s", envValue, expected)
	}
}

func TestGetEnvDoesNotExist(t *testing.T) {
	expected := "foo"
	fallback := expected
	envValue := getEnv("TESTENVKEY", fallback)
	if envValue != expected {
		t.Errorf("Retrun value did not fallback, got: %s, want: %s", envValue, expected)
	}
}
