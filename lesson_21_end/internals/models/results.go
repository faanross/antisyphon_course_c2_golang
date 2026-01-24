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
