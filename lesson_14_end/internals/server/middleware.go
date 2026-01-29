package server

import (
	"log"
	"net/http"
)

// AuthMiddleware returns a middleware that wraps a handler with HMAC authentication
func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := VerifyRequest(r, secret); err != nil {
				log.Printf("Authentication failed: %v", err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
