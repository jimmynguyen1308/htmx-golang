package configs

import (
	"flag"
	"htmx-golang/containers/env"
	"htmx-golang/containers/logger"
	"os"

	"github.com/joho/godotenv"
)

func SetupEnvironmentVariables() {
	devEnv := flag.String("env", ".env.development", "Env file")
	if err := godotenv.Load(*devEnv); err != nil {
		_, exists := os.LookupEnv("ENVIRONMENT")
		if exists {
			return
		}
		logger.Info("No .env file found, creating .env.development...")
		var envFileName string = ".env.development"
		newEnvs := map[string]string{
			"ENVIRONMENT": env.DEVELOPMENT,
			"HOST":        "http://localhost",
			"PORT":        "8080",
		}
		godotenv.Write(newEnvs, envFileName)
		godotenv.Load(envFileName)
	}
}
