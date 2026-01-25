package server

import (
	"fmt"

	"c2framework/internals/config"
)

// Server defines the contract for servers
type Server interface {
	// Start begins listening for requests
	Start() error

	// Stop gracefully shuts down the server
	Stop() error
}

// NewServer creates a new server based on the protocol
func NewServer(cfg *config.ServerConfig) (Server, error) {
	switch cfg.Protocol {
	// TODO implement constructor call for https, nil for error
	case "dns":
		return nil, fmt.Errorf("DNS not yet implemented")
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
