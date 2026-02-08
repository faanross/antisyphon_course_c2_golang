package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"c2framework/internals/config"
)

// HTTPSServer implements the Server interface for HTTPS
type HTTPSServer struct {
	// TODO: create field addr of type string
	// TODO: create server of type *http.Server
	// TODO: add tlsCert of type string
	// TODO: add tlsKey of type string
}

// HTTPSResponse represents the JSON response for HTTPS
type HTTPSResponse struct {
	// TODO: create field Change of type bool, add json tags
}

// NewHTTPSServer creates a new HTTPS server
func NewHTTPSServer(cfg *config.ServerConfig) *HTTPSServer {
	return &HTTPSServer{
		// TODO assign addr by using Sprintf and combining ListeningInterface and ListeningPort
		// TODO: Assign tlsCert as cfg.TlsCert
		// TODO: Assign tlsKey as cfg.TlsKey

	}
}

// Start implements Server.Start for HTTPS
func (s *HTTPSServer) Start() error {
	// Create Chi router
	// TODO: create router r, call NewRouter() from chi library

	// TODO: Define our GET endpoint at /, calls RootHandler

	// Create the HTTP server
	// TODO instantiate s.server as reference to http.Server
	// TODO: Addr is assigned s.addr
	// TODO: Handler is assigned r

	// Start the server
	// TODO call method ListenAndServeTLS, pass cert and key, return call directly
}

// Stop implements Server.Stop for HTTPS
func (s *HTTPSServer) Stop() error {
	// TODO: If there's no server, nothing to stop, return nil

	// TODO: Create context to give the server 5 seconds to shut down gracefully
	defer cancel()

	return s.server.Shutdown(ctx)
}

// RootHandler handles requests to the root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint %s has been hit by agent\n", r.URL.Path)

	// Create response with change set to false
	response := HTTPSResponse{
		// TODO: Assign default to Change as false
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	// TODO: encode response with json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
