package agent

import (
	"context"
	"fmt"

	"c2framework/internals/config"
)

// Agent defines the contract for agents
type Agent interface {
	// Send sends a message and waits for a response
	Send(ctx context.Context) ([]byte, error)
}

// NewAgent creates a new agent based on the protocol
func NewAgent(cfg *config.AgentConfig) (Agent, error) {
	switch cfg.Protocol {
	case "https":
		return NewHTTPSAgent(cfg.ServerIP, cfg.ServerPort), nil
		// TODO: add case for dns, call constructor
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
