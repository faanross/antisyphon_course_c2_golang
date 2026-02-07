//go:build windows

package agent

import (
	"fmt"

	"c2framework/internals/control"
	"c2framework/internals/models"
)

const (
	runKeyPath = `Software\Microsoft\Windows\CurrentVersion\Run`
)

// doPersist handles Registry Run Key persistence on Windows
func doPersist(args control.PersistArgsAgent) models.PersistResult {
	result := models.PersistResult{}

	// Open the Run key
	// TODO: call registry.OpenKey() to open Run Key
	if err != nil {
		result.Success = false
		result.Message = fmt.Sprintf("failed to open registry key: %v", err)
		return result
	}
	// TODO: defer Close() of key

	if args.Remove {
		// Remove the registry value
		// TODO: remove key using key.DeleteValue()
		if err != nil {
			// TODO: set result.Success to false
			result.Success = false
			result.Message = fmt.Sprintf("failed to delete registry value: %v", err)
			return result
		}
		// TODO: set result.Success to true
		result.Message = fmt.Sprintf("Removed registry persistence '%s'", args.Name)
	} else {
		// Set the registry value to our executable path
		// TODO: set run key by calling key.SetStringValue()
		if err != nil {
			result.Success = false
			result.Message = fmt.Sprintf("failed to set registry value: %v", err)
			return result
		}
		result.Success = true
		result.Message = fmt.Sprintf("Installed registry persistence '%s' -> %s", args.Name, args.AgentPath)
	}

	return result
}
