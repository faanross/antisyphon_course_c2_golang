package control

import (
	"encoding/json"
	"fmt"
	"log"
)

// validatePersistCommand validates "persist" command arguments
func validatePersistCommand(rawArgs json.RawMessage) error {
	if len(rawArgs) == 0 {
		return fmt.Errorf("persist command requires arguments")
	}

	var args PersistArgsClient

	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return fmt.Errorf("invalid argument format: %w", err)
	}

	// Validate method
	// TODO: create registry called validMethods with string:bool pairs
	// TODO: add field registry, set to true
	// TODO: add field startup, set to true

	if !validMethods[args.Method] {
		return fmt.Errorf("invalid method '%s' (valid: registry, startup)", args.Method)
	}

	// Name is required
	// TODO: Conditional check if args.Name was provided, if not return error

	log.Printf("Persist validation passed: method=%s, name=%s, remove=%v",
		args.Method, args.Name, args.Remove)
	return nil
}

// processPersistCommand processes persistence arguments
func processPersistCommand(rawArgs json.RawMessage) (json.RawMessage, error) {
	var clientArgs PersistArgsClient

	if err := json.Unmarshal(rawArgs, &clientArgs); err != nil {
		return nil, fmt.Errorf("unmarshaling args: %w", err)
	}

	// Pass through to agent - it knows its own executable path
	agentArgs := PersistArgsAgent{
		// TODO: Set Method equal to field in clientArgs
		// TODO: Set Name equal to field in clientArgs
		// TODO: Set Remove equal to field in clientArgs
		AgentPath: "", // Agent will fill this in
	}

	processedJSON, err := json.Marshal(agentArgs)
	if err != nil {
		return nil, fmt.Errorf("marshaling processed args: %w", err)
	}

	action := "install"

	// TODO: if clientArgs.Remove, then change action to remove

	log.Printf("Persist processed: %s persistence via %s (name: %s)",
		action, clientArgs.Method, clientArgs.Name)
	return processedJSON, nil
}
