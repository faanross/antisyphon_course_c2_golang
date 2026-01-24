package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"c2framework/internals/agent"
	"c2framework/internals/config"
)

func main() {
	// Create agent config - temporarily set to dns for testing
	cfg := &config.AgentConfig{
		Protocol:   "dns",
		ServerIP:   "127.0.0.1",
		ServerPort: "8443",
		Timing: config.TimingConfig{
			Delay:  5 * time.Second,
			Jitter: 50,
		},
	}

	// Call our factory function
	comm, err := agent.NewAgent(cfg)
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Temporary test: single send (runloop doesn't handle DNS yet)
	comm.Send(ctx)

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting down client...")
	cancel()
}
