package control

import "encoding/json"

// CommandClient represents a command with its arguments as sent by Client
type CommandClient struct {
	Command   string          `json:"command"`
	Arguments json.RawMessage `json:"data,omitempty"`
}

// ShellcodeArgsClient contains the command-specific arguments for Shellcode Loader as sent by Client
type ShellcodeArgsClient struct {
	FilePath   string `json:"file_path"`
	ExportName string `json:"export_name"`
}

// ShellcodeArgsAgent - what we send to the agent
// TODO: Create ShellcodeArgsAgent
// TODO: Add field ShellcodeBase64 of type string with json tags
// TODO: Add field ExportName of type string with json tags
