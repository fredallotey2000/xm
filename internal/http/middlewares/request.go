package middleware

import (
	"io"
	"net/http"
)

const (
	maxJsonSize = 20000
)

// RequestBodyLimiter function limits request file size to 2000 bytes
func RequestBodyLimiter(original http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		_, err := io.ReadAll(io.LimitReader(r.Body, maxJsonSize))
		if err != nil {
			http.Error(w, "Not authorized", http.StatusBadRequest)
			return
		} else {
			original.ServeHTTP(w, r)
		}
	})
}
