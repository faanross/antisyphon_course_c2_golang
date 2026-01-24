package shellcode

import "c2framework/internals/models"

// CommandShellcode is the interface for shellcode execution
type CommandShellcode interface {
	DoShellcode(dllBytes []byte, exportName string) (models.ShellcodeResult, error)
}
