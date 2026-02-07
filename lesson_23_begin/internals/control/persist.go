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

	// Name is required
	if args.Name == "" {
		return fmt.Errorf("name is required")
	}

	log.Printf("Persist validation passed: name=%s, remove=%v",
		args.Name, args.Remove)
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
		Name:      clientArgs.Name,
		Remove:    clientArgs.Remove,
		AgentPath: "", // Agent will fill this in
	}

	processedJSON, err := json.Marshal(agentArgs)
	if err != nil {
		return nil, fmt.Errorf("marshaling processed args: %w", err)
	}

	action := "install"
	if clientArgs.Remove {
		action = "remove"
	}
	log.Printf("Persist processed: %s persistence (name: %s)",
		action, clientArgs.Name)
	return processedJSON, nil
}
