package main

import (
	"htmx-golang/configs"
	"htmx-golang/containers/env"
	"htmx-golang/containers/logger"
	"htmx-golang/web"
	"net/http"
	"os"
)

func main() {
	configs.SetupEnvironmentVariables()

	handler := web.Initialize()
	host := env.GetEnv("HOST")
	port := ":" + env.GetEnv("PORT")
	logger.Info("listening on " + host + port)

	if err := http.ListenAndServe(port, handler); err != nil {
		logger.Error("http.ListenAndServe():", "err", err)
		os.Exit(1)
	}
}
