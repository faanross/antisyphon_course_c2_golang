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
	FilePath string `json:"file_path"`
	FileData string `json:"file_data"` // Base64 encoded file contents
	FileSize int64  `json:"file_size"` // Original file size in bytes
	Success  bool   `json:"success"`
	ErrorMsg string `json:"error,omitempty"`
}

// PersistResult - what the agent sends back for persist command
type PersistResult struct {
	// TODO: Add field Method of type string with json tags
	// TODO: Add field Success of type string with json tags
	// TODO: Add field Message of type string with json tags
}
