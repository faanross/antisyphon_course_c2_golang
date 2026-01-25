package agent

import (
	"fmt"

	"c2framework/internals/config"
)

// Agent defines the contract for agents
type Agent interface {
	// Send sends a message and waits for a response
	// TODO: Add signature for Send(), receives context.Context as argument, return byte slice and error

}

// NewAgent creates a new agent based on the protocol
func NewAgent(cfg *config.AgentConfig) (Agent, error) {
	switch cfg.Protocol {
	case "https":
		return nil, fmt.Errorf("HTTPS not yet implemented")
		// TODO: Add a case for dns, similar to https
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
