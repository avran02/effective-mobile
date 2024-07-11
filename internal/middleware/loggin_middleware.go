package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request",
			"method", r.Method,
			"url", r.URL.String(),
			"remote_addr", r.RemoteAddr,
		)
		slog.Debug(fmt.Sprintf("Headers: %+v\nBody: %+v", r.Header, r.Body))

		next.ServeHTTP(w, r)
	})
}
