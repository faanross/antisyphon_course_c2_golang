package control

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// validateShellcodeCommand validates "shellcode" command arguments from client
func validateShellcodeCommand(rawArgs json.RawMessage) error {

	// TODO: if len of rawArgs is 0, means empty, return with Errorf

	// TODO: create args of type ShellcodeArgsClient

	// TODO: unmarshall rawArgs into args, error-check

	// TODO: make sure args.FilePath is not empty, if so return fmt.Errorf

	// TODO: make sure args.ExportName is not empty, if so return fmt.Errorf

	// Check if file exists
	// TODO: Conditional if see if file exists using both os.Stat and os.IsNotExist

	// TODO: Log to terminal validation passed with values

	// TODO: Return nil

}

// processShellcodeCommand reads the DLL file and converts to base64
func processShellcodeCommand(rawArgs json.RawMessage) (json.RawMessage, error) {

	// TODO: Create clientArgs of type ShellcodeArgsClient

	// TODO: unmarshall rawArgs into clientArgs, error-check

	// Read the DLL file
	// TODO: create file by calling os.Open()
	// TODO: error check call

	// TODO: defer file close

	// TODO create fileBytes by calling os.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	// Convert to base64
	// TODO create shellcodeB64 by calling base64.StdEncoding.EncodeToString()

	// Create the arguments that will be sent to the agent
	// TODO: create agentArgs and assign 2 fields with value from above

	// Marshal arguments ready to be sent to agent
	// TODO: create processedJSON by marshalling agentArgs
	// TODO: perform error-check

	// TODO: Print a confirmation message to terminal

	// TODO: return processedJSON and nil

}
