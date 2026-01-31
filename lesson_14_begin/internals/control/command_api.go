package control

import "encoding/json"

// CommandValidator validates command-specific arguments
// TODO: Create CommandValidator type, function, accepts json.RawMessage, returns error

// CommandProcessor processes command-specific arguments
// TODO: Create CommandProcessor type, function, accepts json.RawMessage, returns json.RawMessage + error

// Registry of valid commands with their validators and processors
var validCommands = map[string]struct {
	Validator CommandValidator
	Processor CommandProcessor
}{
	"shellcode": {
		// TODO: assign Validator type validateShellcodeCommand
		// TODO: assign Processor type processShellcodeCommand
	},
	"whoami": {}, // No arguments needed
}
