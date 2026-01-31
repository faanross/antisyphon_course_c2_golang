package models

import "encoding/json"

// ShellcodeResult represents the result of shellcode execution
type ShellcodeResult struct {
	Message string `json:"message"`
}

// AgentTaskResult represents the result of command execution sent back to server
type AgentTaskResult struct {
	JobID         string          `json:"job_id"`
	Success       bool            `json:"success"`
	CommandResult json.RawMessage `json:"command_result,omitempty"`
	Error         string          `json:"error,omitempty"`
}

// DownloadResult - what the agent sends back for download command
type DownloadResult struct {
	// TODO: Add field FilePath of type string + json tags
	// TODO: Add field FileData of type string + json tags
	// TODO: Add field FileSize of type int64 + json tags
	// TODO: Add field Success of type bool + json tags
	// TODO: Add field ErrorMsg of type string + json tags (optional)
	
}
