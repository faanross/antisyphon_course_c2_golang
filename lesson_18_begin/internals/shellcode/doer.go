package shellcode

// ShellcodeResult represents the result of shellcode execution
type ShellcodeResult struct {
	Message string
}

// Doer interface for shellcode execution
type Doer interface {
	DoShellcode(shellcodeBytes []byte, exportName string) (*ShellcodeResult, error)
}
