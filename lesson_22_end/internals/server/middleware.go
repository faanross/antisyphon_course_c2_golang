package server

import (
	"log"
	"net/http"
)

// AuthMiddleware wraps a handler with HMAC authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := VerifyRequest(r); err != nil {
			log.Printf("Authentication failed: %v", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
