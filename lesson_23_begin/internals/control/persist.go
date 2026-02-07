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

	// TODO: create args of type PersistArgsClient

	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return fmt.Errorf("invalid argument format: %w", err)
	}

	// Name is required
	// TODO: Make sure Name is provided

	log.Printf("Persist validation passed: name=%s, remove=%v",
		args.Name, args.Remove)
	return nil
}

// processPersistCommand processes persistence arguments
func processPersistCommand(rawArgs json.RawMessage) (json.RawMessage, error) {

	// TODO: create clientArgs of type PersistArgsClient

	if err := json.Unmarshal(rawArgs, &clientArgs); err != nil {
		return nil, fmt.Errorf("unmarshaling args: %w", err)
	}

	// Pass through to agent - it knows its own executable path
	agentArgs := PersistArgsAgent{
		// TODO: add Name as clientArgs.Name
		// TODO: same for Remove
		// TODO: AgentPath is a blank string for now

	}

	processedJSON, err := json.Marshal(agentArgs)
	if err != nil {
		return nil, fmt.Errorf("marshaling processed args: %w", err)
	}

	// TODO: set action to install
	// TODO if clientArgs.Remove, then invert it to remove

	log.Printf("Persist processed: %s persistence (name: %s)",
		action, clientArgs.Name)
	return processedJSON, nil
}
