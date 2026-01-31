//go:build darwin

package shellcode

import (
	"fmt"

	"c2framework/internals/models"
)

// macShellcode implements the CommandShellcode interface for Darwin/MacOS
// TODO: create new type called macShellcode, no fields

// New is the constructor for our Mac-specific Shellcode command
func New() CommandShellcode {
	// TODO return instantiated macShellcode
}

// DoShellcode is the stub implementation for macOS
func (ms *macShellcode) DoShellcode(dllBytes []byte, exportName string) (models.ShellcodeResult, error) {
	fmt.Println("|SHELLCODE DOER MACOS| This feature has not yet been implemented for MacOS.")

	// TODO: create result of type models.ShellcodeResult, for Message field simply write a failure message

	return result, nil
}
