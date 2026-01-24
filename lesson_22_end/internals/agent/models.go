package agent

import (
	"context"
	"encoding/json"
	"fmt"

	"c2framework/internals/config"
	"c2framework/internals/models"
)

// Agent defines the contract for agents
type Agent interface {
	// Send sends a message and waits for a response
	Send(ctx context.Context) (json.RawMessage, error)
}

// AgentTaskResult is an alias to the shared models type
type AgentTaskResult = models.AgentTaskResult

// NewAgent creates a new agent based on the protocol
func NewAgent(cfg *config.AgentConfig) (Agent, error) {
	switch cfg.Protocol {
	case "https":
		return NewHTTPSAgent(cfg.ServerIP, cfg.ServerPort), nil
	case "dns":
		return NewDNSAgent(cfg.ServerIP, cfg.ServerPort), nil
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
