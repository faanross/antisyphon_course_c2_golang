package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"c2framework/internals/config"
	"c2framework/internals/control"
	"c2framework/internals/crypto"
)

// HTTPSServer implements the Server interface for HTTPS
type HTTPSServer struct {
	addr         string
	server       *http.Server
	tlsCert      string
	tlsKey       string
	sharedSecret string
}

// HTTPSResponse represents the JSON response for HTTPS
type HTTPSResponse struct {
	Change bool `json:"change"`
}

// NewHTTPSServer creates a new HTTPS server
func NewHTTPSServer(cfg *config.ServerConfig) *HTTPSServer {
	return &HTTPSServer{
		addr:         fmt.Sprintf("%s:%s", cfg.ListeningInterface, cfg.ListeningPort),
		tlsCert:      cfg.TlsCert,
		tlsKey:       cfg.TlsKey,
		sharedSecret: cfg.SharedSecret,
	}
}

// Start implements Server.Start for HTTPS
func (s *HTTPSServer) Start() error {
	// Create Chi router
	r := chi.NewRouter()

	// Apply authentication middleware to agent routes
	r.With(AuthMiddleware(s.sharedSecret)).Post("/", RootHandler(s.sharedSecret))

	// Create the HTTP server
	s.server = &http.Server{
		Addr:    s.addr,
		Handler: r,
	}

	// Start the server
	return s.server.ListenAndServeTLS(s.tlsCert, s.tlsKey)
}

// Stop implements Server.Stop for HTTPS
func (s *HTTPSServer) Stop() error {
	// If there's no server, nothing to stop
	if s.server == nil {
		return nil
	}

	// Give the server 5 seconds to shut down gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.server.Shutdown(ctx)
}

// RootHandler returns a handler that encrypts responses
func RootHandler(secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Endpoint %s has been hit by agent\n", r.URL.Path)

		// Read encrypted body
		encryptedBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusBadRequest)
			return
		}

		log.Printf("Payload pre-decryption: %s", string(encryptedBody))

		// Decrypt the payload
		// TODO: call crypto.Decrypt() to decrypt, assign return to new car plaintext
		if err != nil {
			log.Printf("Decryption failed: %v", err)
			http.Error(w, "Decryption failed", http.StatusBadRequest)
			return
		}

		log.Printf("Payload post-decryption: %s", string(plaintext))

		// Check if we should transition
		shouldChange := control.Manager.CheckAndReset()
		response := HTTPSResponse{
			Change: shouldChange,
		}
		if shouldChange {
			log.Printf("HTTPS: Sending transition signal (change=true)")
		} else {
			log.Printf("HTTPS: Normal response (change=false)")
		}

		// Marshal response to JSON
		responseJSON, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error marshaling response: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Encrypt the response
		encryptedResponse, err := crypto.Encrypt(responseJSON, secret)
		if err != nil {
			log.Printf("Error encrypting response: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set content type to octet-stream for encrypted data
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write([]byte(encryptedResponse))
	}
}
