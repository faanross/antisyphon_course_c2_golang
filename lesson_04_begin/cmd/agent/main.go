package main

import (
	"c2framework/internals/agent"
	"c2framework/internals/config"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Create agent config directly in code
	cfg := &config.AgentConfig{
		Protocol:   "https",
		ServerIP:   "127.0.0.1",
		ServerPort: "8443",
		Timing: config.TimingConfig{
			// TODO: Assign Delay as 5 seconds
			// TODO: Assign Jitter as 50 (as in 50%)

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

	// Start run loop in goroutine
	go func() {
		log.Printf("Starting %s client run loop", cfg.Protocol)
		log.Printf("Delay: %v, Jitter: %d%%", cfg.Timing.Delay, cfg.Timing.Jitter)

		// TODO now call agent.RunLoop to start our server
		
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting down client...")
	cancel() // This will cause the run loop to exit
}
