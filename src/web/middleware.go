package web

import (
	"context"
	"fmt"
	"htmx-golang/containers/logger"
	"net/http"
)

func (mw New) tracing() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = fmt.Sprintf("%d", mw.Time.Now().UnixNano())
			}
			ctx := context.WithValue(r.Context(), mw.ReqIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func logging(reqIDKey key) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(reqIDKey).(string)
				if !ok {
					requestID = "unknown"
					logger.Warn(
						r.URL.Path,
						"reqID", requestID,
						"method", r.Method,
						"remoteAddr", r.RemoteAddr,
						"userAgent", r.UserAgent(),
					)
				}
				logger.Info(
					r.URL.Path,
					"reqID", requestID,
					"method", r.Method,
					"remoteAddr", r.RemoteAddr,
					"userAgent", r.UserAgent(),
				)
			}()
			next.ServeHTTP(w, r)
		})
	}
}
