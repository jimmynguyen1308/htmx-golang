package web

import (
	"htmx-golang/containers/time"
	"net/http"
)

type key int

type New struct {
	ReqIDKey key
	Time     time.Time
}

var Default = New{
	ReqIDKey: 0,
	Time:     time.RealTime{},
}

func (w New) generateWebHandler(router http.Handler) http.Handler {
	return w.tracing()(logging(w.ReqIDKey)(router))
}
