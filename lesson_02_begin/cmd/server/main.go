package main

import (
	"c2framework/internals/config"
	"log"
)

func main() {
	// Create server config directly in code
	cfg := &config.ServerConfig{
		// TODO: Assign protocol https
		// TODO: Assign ListeningInterface as 0.0.0.0
		// TODO: Assign ListeningPort as 8443
		// TODO: Assign TlsCert path
		// TODO: Assign TlsKey path
	}

	// TODO: Create server using interface's factory function
	// TODO: Error-check

	// Start the server in own goroutine
	// TODO: create a goroutine
	log.Printf("Starting %s server on %s:%s", cfg.Protocol, cfg.ListeningInterface, cfg.ListeningPort)

	// TODO: call srv.Start()
	// TODO: Error-check

	// Wait for interrupt signal
	// TODO: Assign sigChan using make, channel type is os.Signal, buffer of 1)
	// TODO: call signal.Notify, pass sigChan and os.Interrupt
	// TODO: Block main goroutine here with sigChan

	// Graceful shutdown
	log.Println("Shutting down server...")

	//TODO: call srv.Stop(), and error check

}
