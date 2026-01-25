package main

import (
	"log"
	"os"
	"os/signal"

	"c2framework/internals/config"
	"c2framework/internals/server"
)

func main() {
	// Create server config directly in code
	cfg := &config.ServerConfig{
		// TODO: Assign protocol https
		// TODO: Assign ListeningInterface as 0.0.0.0
		// TODO: Assign ListeningPort as 8443
		TlsCert: "./certs/server.crt",
		TlsKey:  "./certs/server.key",
	}

	// TODO: Create server using interface's factory function
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Start the server in own goroutine
	go func() {
		log.Printf("Starting %s server on %s:%s", cfg.Protocol, cfg.ListeningInterface, cfg.ListeningPort)
		if err := srv.Start(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	// TODO: Assign sigChan using make, channel type is os.Signal, buffer of 1)
	signal.Notify(sigChan, os.Interrupt)
	// TODO: Block main goroutine here with sigChan

	// Graceful shutdown
	log.Println("Shutting down server...")
	if err := srv.Stop(); err != nil {
		log.Printf("Error stopping server: %v", err)
	}
}
