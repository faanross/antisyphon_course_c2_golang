package server

import (
	"fmt"

	"c2framework/internals/config"
)

// Server defines the contract for servers
type Server interface {
	// TODO: Add Start(), no arguments, returns error
	// Start begins listening for requests

	// TODO: Add Stop(), no arguments, returns error
	// Stop gracefully shuts down the server

}

// NewServer creates a new server based on the protocol
func NewServer(cfg *config.ServerConfig) (Server, error) {
	switch cfg.Protocol {
	// TODO: Add a case for https, similar to dns
	case "dns":
		return nil, fmt.Errorf("DNS not yet implemented")
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
