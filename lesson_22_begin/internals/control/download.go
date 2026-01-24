package control

import (
	"encoding/json"
	"fmt"
	"log"
)

// validateDownloadCommand validates "download" command arguments from client
func validateDownloadCommand(rawArgs json.RawMessage) error {
	if len(rawArgs) == 0 {
		return fmt.Errorf("download command requires arguments")
	}

	var args DownloadArgsClient

	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return fmt.Errorf("invalid argument format: %w", err)
	}

	if args.FilePath == "" {
		return fmt.Errorf("file_path is required")
	}

	log.Printf("Download validation passed: file_path=%s", args.FilePath)
	return nil
}

// processDownloadCommand processes download arguments (minimal for this command)
func processDownloadCommand(rawArgs json.RawMessage) (json.RawMessage, error) {
	var clientArgs DownloadArgsClient

	if err := json.Unmarshal(rawArgs, &clientArgs); err != nil {
		return nil, fmt.Errorf("unmarshaling args: %w", err)
	}

	// For download, we just pass the file path as-is
	agentArgs := DownloadArgsAgent{
		FilePath: clientArgs.FilePath,
	}

	processedJSON, err := json.Marshal(agentArgs)
	if err != nil {
		return nil, fmt.Errorf("marshaling processed args: %w", err)
	}

	log.Printf("Download processed: requesting file %s from agent", clientArgs.FilePath)
	return processedJSON, nil
}
