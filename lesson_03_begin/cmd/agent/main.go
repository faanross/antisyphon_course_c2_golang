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

	// TODO: call NewAgent() constructor
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	// TEMPORARY CODE JUST TO TEST!
	// Send a test message

	log.Printf("Sending request to %s server...", cfg.Protocol)
	// TODO: call Send() method on a, pass context.Background() as argument
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Parse and display response
	// TODO: instantiate httpsResp of type server.HTTPSResponse to unmarshall into
	// TODO: unmarshall into struct + error-check

	// TODO use log.Printf to print httpsResp.Change to stdout
}
