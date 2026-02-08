package agent

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

// HTTPSAgent implements the Agent interface for HTTPS
type HTTPSAgent struct {
	// TODO: add serverAddr of type string
	// TODO: add client of type *http.Client
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
			// TODO: assign TLSClientConfig as struct we created above
		},
	}

	return &HTTPSAgent{
		// TODO: assign serverAddr serverIP:serverPort using fmt.Sprintf
		// TODO: Assign client as client (from above)
	}
}

// Send implements Agent.Send for HTTPS
func (c *HTTPSAgent) Send(ctx context.Context) ([]byte, error) {
	// TODO: Construct url with Sprintf and serverAddr

	// Create GET request
	// TODO: Call http.NewRequestWithContext() to great request
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Send request
	// TODO: call c.client.Do(), pass req as argument, to send our request
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("server returned status %d: %s", resp.StatusCode, body)
	}

	// Read response body
	// TODO: read the response body using io.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	// Return the raw JSON as message data
	return body, nil
}
