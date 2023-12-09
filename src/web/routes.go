package web

import (
	"net/http"

	"github.com/jritsema/gotoolbox/web"
)

func handleRoutes() {
	router.Handle("/css/output.css", http.FileServer(http.FS(css)))

	router.Handle("/", web.Action(index))
	router.Handle("/index.html", web.Action(index))
}

func index(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "index.html", nil, nil)
}
