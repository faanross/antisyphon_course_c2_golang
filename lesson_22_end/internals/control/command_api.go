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
	"download": {
		Validator: validateDownloadCommand,
		Processor: processDownloadCommand,
	},
	"whoami": {}, // No arguments needed
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
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if len(cq.PendingCommands) == 0 {
		return CommandClient{}, false
	}

	cmd := cq.PendingCommands[0]
	cq.PendingCommands = cq.PendingCommands[1:]

	log.Printf("DEQUEUED: Command '%s'", cmd.Command)

	return cmd, true
}
