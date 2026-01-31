package agent

import (
	"context"
	"encoding/json"
	"fmt"

	"c2framework/internals/config"
)

// Agent defines the contract for agents
type Agent interface {
	// Send sends a message and waits for a response
	Send(ctx context.Context) (json.RawMessage, error)
}

// AgentTaskResult represents the result of command execution sent back to server
type AgentTaskResult struct {
	// TODO: Add field JobID of type string, json tags
	// TODO: Add field Success of type bool, json tags
	// TODO: Add field CommandResult of type json.RawMessage, json tags (optional)
	// TODO: Add field Error of type string, json tags (optional)

}

// NewAgent creates a new agent based on the protocol
func NewAgent(cfg *config.AgentConfig) (Agent, error) {
	switch cfg.Protocol {
	case "https":
		return NewHTTPSAgent(cfg.ServerIP, cfg.ServerPort, cfg.SharedSecret), nil
	case "dns":
		return NewDNSAgent(cfg.ServerIP, cfg.ServerPort), nil
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
