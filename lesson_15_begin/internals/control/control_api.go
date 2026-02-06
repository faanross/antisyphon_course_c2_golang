package control

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/go-chi/chi/v5"
)

// TransitionManager handles the global transition state
type TransitionManager struct {
	mu               sync.RWMutex
	shouldTransition bool
}

// Global instance
var Manager = &TransitionManager{
	shouldTransition: false,
}

// TriggerTransition sets the transition flag
func (tm *TransitionManager) TriggerTransition() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.shouldTransition = true
	log.Printf("Transition triggered")
}

// CheckAndReset atomically checks if transition is needed and resets the flag
// This ensures the transition signal is consumed only once
func (tm *TransitionManager) CheckAndReset() bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.shouldTransition {
		tm.shouldTransition = false // Reset immediately
		log.Printf("Transition signal consumed and reset")
		return true
	}

	return false
}

// StartControlAPI starts the control API server on port 8080
func StartControlAPI() {
	// Create Chi router
	r := chi.NewRouter()

	r.Post("/switch", handleSwitch)
	r.Post("/command", commandHandler)

	log.Println("Starting Control API on :8080")
	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Printf("Control API error: %v", err)
		}
	}()
}

func handleSwitch(w http.ResponseWriter, r *http.Request) {
	Manager.TriggerTransition()

	response := "Protocol transition triggered"

	json.NewEncoder(w).Encode(response)
}

func commandHandler(w http.ResponseWriter, r *http.Request) {

	// Instantiate custom type to receive command from client
	var cmdClient CommandClient

	// Unmarshal the request body
	if err := json.NewDecoder(r.Body).Decode(&cmdClient); err != nil {
		log.Printf("ERROR: Failed to decode JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error decoding JSON")
		return
	}

	// Normalize command to lowercase
	cmdClient.Command = strings.ToLower(cmdClient.Command)

	var commandReceived = fmt.Sprintf("Received command: %s", cmdClient.Command)
	log.Printf(commandReceived)

	// STEP 1: Check if command exists
	cmdConfig, exists := validCommands[cmdClient.Command]
	if !exists {
		var commandInvalid = fmt.Sprintf("ERROR: Unknown command: %s", cmdClient.Command)
		log.Printf(commandInvalid)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commandInvalid)
		return
	}

	// STEP 2: Validate arguments (if validator exists)
	if cmdConfig.Validator != nil {
		if err := cmdConfig.Validator(cmdClient.Arguments); err != nil {
			var commandInvalid = fmt.Sprintf("ERROR: Validation failed for '%s': %v", cmdClient.Command, err)
			log.Printf(commandInvalid)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(commandInvalid)
			return
		}
	}

	// STEP 3: Process arguments (if processor exists)
	if cmdConfig.Processor != nil {
		processedArgs, err := cmdConfig.Processor(cmdClient.Arguments)
		if err != nil {
			var commandInvalid = fmt.Sprintf("ERROR: Processing failed for '%s': %v", cmdClient.Command, err)
			log.Printf(commandInvalid)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(commandInvalid)
			return
		}
		cmdClient.Arguments = processedArgs
		log.Printf("Processed command arguments: %s", cmdClient.Command)
	}

	// Queue the validated and processed command
	// TODO call addCommand() method on AgentCommands to queue command

	// Confirm on the client side command was received
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commandReceived)
}
