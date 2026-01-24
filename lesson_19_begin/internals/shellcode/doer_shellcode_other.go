//go:build !darwin && !windows

package shellcode

import (
	"c2framework/internals/models"
)

// otherShellcode implements the CommandShellcode interface for other platforms
type otherShellcode struct{}

// New is the constructor for other platforms
func New() CommandShellcode {
	return &otherShellcode{}
}

// DoShellcode is the stub implementation for unsupported platforms
func (os *otherShellcode) DoShellcode(dllBytes []byte, exportName string) (models.ShellcodeResult, error) {
	return models.ShellcodeResult{
		Message: "Shellcode execution not implemented for this platform",
	}, nil
}
