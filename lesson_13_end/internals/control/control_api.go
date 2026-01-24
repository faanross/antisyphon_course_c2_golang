package control

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	// The first thing we need to do is unmarshal the request body into the custom type
	if err := json.NewDecoder(r.Body).Decode(&cmdClient); err != nil {
		log.Printf("ERROR: Failed to decode JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error decoding JSON")
		return
	}

	// Visually confirm we get the command we expected
	var commandReceived = fmt.Sprintf("Received command: %s", cmdClient.Command)
	log.Printf(commandReceived)

	// Confirm on the client side command was received
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commandReceived)
}
