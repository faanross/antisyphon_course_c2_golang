package main

import (
	"context"
	"encoding/json"
	"log"

	"c2framework/internals/agent"
	"c2framework/internals/config"
	"c2framework/internals/server"
)

func main() {
	// Create agent config directly in code
	cfg := &config.AgentConfig{
		Protocol:   "https",
		ServerIP:   "127.0.0.1",
		ServerPort: "8443",
	}

	a, err := agent.NewAgent(cfg)
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	// TEMPORARY CODE JUST TO TEST!
	// Send a test message

	log.Printf("Sending request to %s server...", cfg.Protocol)
	response, err := a.Send(context.Background())
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Parse and display response
	var httpsResp server.HTTPSResponse
	if err := json.Unmarshal(response, &httpsResp); err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}

	log.Printf("Received response: change=%v", httpsResp.Change)
}
