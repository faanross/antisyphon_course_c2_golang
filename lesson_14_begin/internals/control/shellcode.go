package control

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// validateShellcodeCommand validates "shellcode" command arguments from client
func validateShellcodeCommand(rawArgs json.RawMessage) error {
	if len(rawArgs) == 0 {
		return fmt.Errorf("shellcode command requires arguments")
	}

	// TODO: create args of type ShellcodeArgsClient

	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return fmt.Errorf("invalid argument format: %w", err)
	}

	// TODO: make sure args.FilePath is not empty, if so return fmt.Errorf

	if args.ExportName == "" {
		return fmt.Errorf("export_name is required")
	}

	// Check if file exists
	// TODO: Conditional if see if file exists using both os.Stat and os.IsNotExist

	{
		return fmt.Errorf("file does not exist: %s", args.FilePath)
	}

	log.Printf("Validation passed: file_path=%s, export_name=%s", args.FilePath, args.ExportName)

	return nil
}

// processShellcodeCommand reads the DLL file and converts to base64
func processShellcodeCommand(rawArgs json.RawMessage) (json.RawMessage, error) {

	// TODO: Create clientArgs of type ShellcodeArgsClient

	if err := json.Unmarshal(rawArgs, &clientArgs); err != nil {
		return nil, fmt.Errorf("unmarshaling args: %w", err)
	}

	// Read the DLL file
	// TODO create file by calling os.Open()

	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	// TODO create fileBytes by calling os.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	// Convert to base64
	// TODO create shellcodeB64 by calling base64.StdEncoding.EncodeToString()
	
	// Create the arguments that will be sent to the agent
	agentArgs := ShellcodeArgsAgent{
		ShellcodeBase64: shellcodeB64,
		ExportName:      clientArgs.ExportName,
	}

	// Marshal arguments ready to be sent to agent
	processedJSON, err := json.Marshal(agentArgs)
	if err != nil {
		return nil, fmt.Errorf("marshaling processed args: %w", err)
	}

	log.Printf("Processed file: %s (%d bytes) -> base64 (%d chars)",
		clientArgs.FilePath, len(fileBytes), len(shellcodeB64))

	return processedJSON, nil
}
