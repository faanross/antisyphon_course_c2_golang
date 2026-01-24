package shellcode

// ShellcodeResult represents the result of shellcode execution
type ShellcodeResult struct {
	Message string
}

// Doer interface for shellcode execution
type Doer interface {
	DoShellcode(shellcodeBytes []byte, exportName string) (*ShellcodeResult, error)
}

// ShellcodeDoer is a stub implementation
type ShellcodeDoer struct{}

// New creates a new shellcode doer
func New() *ShellcodeDoer {
	return &ShellcodeDoer{}
}

// DoShellcode is a stub that will be implemented per-OS in future lessons
func (s *ShellcodeDoer) DoShellcode(shellcodeBytes []byte, exportName string) (*ShellcodeResult, error) {
	// Stub implementation - actual implementation will be OS-specific
	return &ShellcodeResult{
		Message: "Shellcode execution not implemented for this platform",
	}, nil
}
