package env

import "os"

const (
	DEVELOPMENT = "development"
	TESTING     = "testing"
	STAGING     = "staging"
	PRODUCTION  = "production"
)

func GetEnv(name string) string {
	return os.Getenv(name)
}

func IsEnv(name string, envs ...string) bool {
	for _, e := range envs {
		if GetEnv(name) == e {
			return true
		}
	}
	return false
}
