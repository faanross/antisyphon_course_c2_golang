package control

import (
	"encoding/json"
	"fmt"
	"log"
)

// validateDownloadCommand validates "download" command arguments from client
func validateDownloadCommand(rawArgs json.RawMessage) error {

	// TODO: Check to ensure rawArgs len is not 0, if it is return an error

	// TODO: Create args of type DownloadArgs

	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return fmt.Errorf("invalid argument format: %w", err)
	}

	// TODO: If args.FilePath is empty return an error

	log.Printf("Download validation passed: file_path=%s", args.FilePath)
	return nil
}
