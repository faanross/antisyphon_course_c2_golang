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
type ShellcodeArgsAgent struct {
	ShellcodeBase64 string `json:"shellcode_base64"`
	ExportName      string `json:"export_name"`
}

// DownloadArgs - arguments for download command (no transformation needed)
type DownloadArgs struct {
	FilePath string `json:"file_path"` // Path on agent's machine
}

// PersistArgsClient - what the client sends
type PersistArgsClient struct {
	// TODO add field Method of type string, json tags, indicates registry or startup
	// TODO add field Name of type string, json tags
	// TODO add field Remove of type bool, json tags
}

// PersistArgsAgent - what we send to the agent
type PersistArgsAgent struct {
	Method string `json:"method"`
	Name   string `json:"name"`
	Remove bool   `json:"remove"`
	// TODO: add field AgentPath of type string with json tags
}
