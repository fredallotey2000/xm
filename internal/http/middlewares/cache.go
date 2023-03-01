package middleware

import (
	"net/http"
)

// JSONMiddleware function to set all response header to Content-Type application/json
func SetCacheControl(cc string,next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", cc)
		next.ServeHTTP(w, r)
	})
}
