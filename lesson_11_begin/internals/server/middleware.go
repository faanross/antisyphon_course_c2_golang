package server

import (
	"net/http"
)

// AuthMiddleware returns a middleware that wraps a handler with HMAC authentication
func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Call VerifyRequest()
			// TODO: if failed: log that request was unauthorized

			// TODO: call next.ServeHTTP(w, r)
			
		})
	}
}
