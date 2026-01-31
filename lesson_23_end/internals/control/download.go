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

	var args DownloadArgs

	if err := json.Unmarshal(rawArgs, &args); err != nil {
		return fmt.Errorf("invalid argument format: %w", err)
	}

	if args.FilePath == "" {
		return fmt.Errorf("file_path is required")
	}

	log.Printf("Download validation passed: file_path=%s", args.FilePath)
	return nil
}
