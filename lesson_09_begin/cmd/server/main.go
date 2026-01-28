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

	// Load our control API
	control.StartControlAPI()

	// Create BOTH servers regardless of config
	log.Printf("Starting both protocol servers on %s:%s", cfg.ListeningInterface, cfg.ListeningPort)

	// Create HTTPS server
	// TODO create httpsCfg as a pointer to cfg
	// TODO manually set protocol field to https
	// TODO create httpsServer by calling server.NewServer, pass reference to httpsCfg

	if err != nil {
		log.Fatalf("Failed to create HTTPS server: %v", err)
	}

	// Create DNS server
	// TODO create dnsCfg as a pointer to cfg
	// TODO manually set protocol field to dns
	// TODO create dnsServer by calling server.NewServer, pass reference to dnsCfg

	if err != nil {
		log.Fatalf("Failed to create DNS server: %v", err)
	}

	// Start HTTPS server in goroutine
	go func() {
		log.Printf("Starting HTTPS server on %s:%s (TCP)", cfg.ListeningInterface, cfg.ListeningPort)
		if err := httpsServer.Start(); err != nil {
			log.Fatalf("HTTPS server error: %v", err)
		}
	}()

	// Start DNS server in goroutine
	// TODO: Start() dnsServer in exactly same way as above, in its own goroutine

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	// Graceful shutdown
	log.Println("Shutting down both servers...")

	if err := httpsServer.Stop(); err != nil {
		log.Printf("Error HTTPS stopping server: %v", err)
	}

	if err := dnsServer.Stop(); err != nil {
		log.Printf("Error DNS stopping server: %v", err)
	}
}
