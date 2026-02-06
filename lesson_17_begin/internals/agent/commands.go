package agent

import (
	"encoding/json"
	"log"

	"c2framework/internals/server"
)

// OrchestratorFunc defines the signature for command orchestrator functions
// TODO: define function type OrchestratorFunc, arguments are pointers to HTTPSAgent and server.HTTPSResponse, returns AgentTaskResult

// registerCommands registers all available command orchestrators
// TODO: create function registerCommands, pointer to HTTPSAgent as arg
// TODO: register "shellcode" agent.commandOrchestrators[], assign value as (*HTTPSAgent).orchestrateShellcode
// Register other commands here in the future

// ExecuteTask receives a command from the server and routes it to the appropriate orchestrator
func (agent *HTTPSAgent) ExecuteTask(job *server.HTTPSResponse) {
	log.Printf("AGENT IS NOW PROCESSING COMMAND %s with ID %s", job.Command, job.JobID)

	// TODO: create result of type AgentTaskResult

	// Look up the orchestrator for this command
	// TODO: look up job.Command key in agent.commandOrchestrators[], assign value and found bool

	if found {
		// Call the orchestrator
		// TODO: call orchestrator with agent and job args, assign return to result
	} else {
		// Command not recognized
		log.Printf("|WARN AGENT TASK| Received unknown command: '%s' (ID: %s)", job.Command, job.JobID)

		// TODO: create result, of type AgentTaskResult
		// TODO: JobID = job.JobID
		// TODO: Success is false in this case
		// TODO: Error something like "command not found"

	}

	// Marshal the result before sending it back
	// TODO marshall result, save as resultBytes, error check

	log.Printf("|AGENT TASK|-> Sending result for Task ID %s (%d bytes)...", job.JobID, len(resultBytes))

	// Send the result back to the server
	// TODO: call agent.SendResult(), pass resultBytes as arg, error-check

	log.Printf("|AGENT TASK|-> Successfully sent result for Task ID %s.", job.JobID)
}
