package env_test

import (
	"htmx-golang/containers/env"
	"testing"
)

func TestGetEnv(t *testing.T) {
	t.Setenv("ENVIRONMENT", env.TESTING)
	if env.GetEnv("ENVIRONMENT") != env.TESTING {
		t.Errorf("Expected %s, got %s: incorrect environment variable", env.TESTING, env.GetEnv("ENVIRONMENT"))
	}
	if env.GetEnv("RANDOM_NAME") != "" {
		t.Errorf("Expected %s, got %s: incorrect environment variable", "", env.GetEnv("RANDOM_NAME"))
	}
}

func TestIsEnv(t *testing.T) {
	t.Setenv("ENVIRONMENT", env.DEVELOPMENT)
	if !env.IsEnv("ENVIRONMENT", env.DEVELOPMENT) {
		t.Errorf("Expected IsEnv(%s) to be true, but got false", env.DEVELOPMENT)
	}
	if env.IsEnv(env.PRODUCTION) {
		t.Errorf("Expected IsEnv(%s) to be false, but got true", env.PRODUCTION)
	}
}
