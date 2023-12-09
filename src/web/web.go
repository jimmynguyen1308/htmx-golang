package web

import (
	"embed"
	"fmt"
	"html/template"
	"htmx-golang/containers/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jritsema/gotoolbox/web"
)

var (
	//go:embed all:templates/*
	templateFS embed.FS

	//go:embed css/output.css
	css embed.FS

	router = http.NewServeMux()

	html *template.Template

	err error
)

func Initialize() http.Handler {
	handleSigTerms()

	html, err = web.TemplateParseFSRecursive(templateFS, ".html", true, nil)
	if err != nil {
		logger.Fatal("web.TemplateParseFSRecursive():", "err", err)
	}

	handleRoutes()

	return Default.generateWebHandler(router)
}

func handleSigTerms() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("received SIGTERM, exiting...")
		os.Exit(1)
	}()
}
