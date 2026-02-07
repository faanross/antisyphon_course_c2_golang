package agent

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"c2framework/internals/control"
	"c2framework/internals/server"
	"c2framework/internals/shellcode"
)

// orchestrateShellcode is the orchestrator for the "shellcode" command
func (agent *HTTPSAgent) orchestrateShellcode(job *server.HTTPSResponse) AgentTaskResult {

	// Create an instance of the shellcode args struct
	// TODO: create shellcodeArgs of type control.ShellcodeArgsAgent

	// ServerResponse.Arguments contains the command-specific args, so now we unmarshal the field into the struct
	// TODO: Unmarshal job.Arguments into shellcodeArgs
	// TODO: Error-check, return AgentTaskResult if failed

	log.Printf("|SHELLCODE ORCHESTRATOR| Task ID: %s. Executing Shellcode, Export Function: %s, ShellcodeLen(b64)=%d\n",
		job.JobID, shellcodeArgs.ExportName, len(shellcodeArgs.ShellcodeBase64))

	// Some basic agent-side validation
	// TODO: check if shellcodeArgs.ShellcodeBase64 is empty

	// TODO: Validate that shellcodeArgs.ExportName is empty

	// Now let's decode our b64
	// TODO: call base64.StdEncoding.DecodeString to decode shellcodeArgs.ShellcodeBase64 as rawShellcode

	if err != nil {
		log.Printf("|ERR SHELLCODE ORCHESTRATOR| Task ID %s: Failed to decode ShellcodeBase64: %v", job.JobID, err)
		return AgentTaskResult{
			JobID:   job.JobID,
			Success: false,
			Error:   "Failed to decode shellcode",
		}
	}

	// Call the "doer" function
	// TODO: call constructor via interface - shellcode.New() - assign return as commandShellcode

	// TODO: call DoShellCode() on commandShellcode

	finalResult := AgentTaskResult{
		JobID: job.JobID,
	}

	// TODO: Marshall shellcodeResult.Message, save as outputJSON

	// TODO: set finalResult.CommandResult as outputJSON

	if err != nil {
		loaderError := fmt.Sprintf("|ERR SHELLCODE ORCHESTRATOR| Loader execution error for TaskID %s: %v. Loader Message: %s",
			job.JobID, err, shellcodeResult.Message)
		// TODO: print loaderError to terminal
		// TODO: set finalResult.Error equal to loaderError
		// TODO: finalResult.Success equal to false

	} else {
		log.Printf("|SHELLCODE SUCCESS| Shellcode execution initiated successfully for TaskID %s. Loader Message: %s",
			job.JobID, shellcodeResult.Message)
		// TODO: finalResult.Success equal to true
	}

	// TODO: return finalResult
}
