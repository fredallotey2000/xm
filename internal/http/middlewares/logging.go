package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.WithFields(
			log.Fields{
				"remote address": r.RemoteAddr,
				"method":         r.Method,
				"path":           r.URL.Path,
				//"status":         w.,
				"latency_ns": time.Since(start).Microseconds(),
			},
		).Info("handle request")
		next.ServeHTTP(w, r)
	})
}
