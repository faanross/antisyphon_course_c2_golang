package control

import "encoding/json"

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
