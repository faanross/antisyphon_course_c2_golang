package main

import (
	"log"
	"os"
	"os/signal"

	"c2framework/internals/config"
	"c2framework/internals/control"
	"c2framework/internals/server"
)

func main() {
	// Create server config directly in code
	cfg := &config.ServerConfig{
		Protocol:           "https",
		ListeningInterface: "127.0.0.1",
		ListeningPort:      "8443",
		TlsCert:            "./certs/server.crt",
		TlsKey:             "./certs/server.key",
	}

	// Start our control API
	control.StartControlAPI()

	// Create server using interface's factory function
	srv, err := server.NewServer(cfg)
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
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	// Graceful shutdown
	log.Println("Shutting down server...")
	if err := srv.Stop(); err != nil {
		log.Printf("Error stopping server: %v", err)
	}
}
