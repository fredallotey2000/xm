package middleware

import (
	"context"
	"net/http"
	"time"
)

const (
	RequestTimeout = 15 * time.Second
)

// TimeoutMiddleware function times out all requests after
func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), RequestTimeout)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
