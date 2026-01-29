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
		SharedSecret:       "your-super-secret-key-change-in-production",
	}

	// Load our control API
	control.StartControlAPI()

	// Create BOTH servers regardless of config
	log.Printf("Starting both protocol servers on %s:%s", cfg.ListeningInterface, cfg.ListeningPort)

	// Create HTTPS server
	httpsCfg := *cfg
	httpsCfg.Protocol = "https"
	httpsServer, err := server.NewServer(&httpsCfg)
	if err != nil {
		log.Fatalf("Failed to create HTTPS server: %v", err)
	}

	// Create DNS server
	dnsCfg := *cfg
	dnsCfg.Protocol = "dns"
	dnsServer, err := server.NewServer(&dnsCfg)
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
	go func() {
		log.Printf("Starting DNS server on %s:%s (UDP)", cfg.ListeningInterface, cfg.ListeningPort)
		if err := dnsServer.Start(); err != nil {
			log.Fatalf("DNS server error: %v", err)
		}
	}()

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
