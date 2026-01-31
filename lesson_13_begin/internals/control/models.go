package control

// CommandClient represents a command with its arguments as sent by Client
type CommandClient struct {
	// TODO: Add field Command of type string, json tags
	// TODO: Add field Arguments of type json.RawMessage, json tags (optional)

}

// ShellcodeArgsClient contains the command-specific arguments for Shellcode Loader as sent by Client
type ShellcodeArgsClient struct {
	// TODO: Add field FilePath of type string, json tags
	// TODO: Add field ExportName of type string, json tags
}
