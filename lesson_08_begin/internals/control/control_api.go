package control

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// TransitionManager handles the global transition state
type TransitionManager struct {
	// TODO: create RW mutex field
	// TODO: create field shouldTransition of type bool
}

// Global instance
var Manager = &TransitionManager{
	// TODO: set shouldTransition to false
}

// TriggerTransition sets the transition flag
func (tm *TransitionManager) TriggerTransition() {

	// TODO: lock the mutex
	// TODO: defer mutex unlock

	// TODO: set shouldTransition equal to true

	log.Printf("Transition triggered")
}

// CheckAndReset atomically checks if transition is needed and resets the flag
// This ensures the transition signal is consumed only once
func (tm *TransitionManager) CheckAndReset() bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.shouldTransition {
		// TODO set tm.shouldTransition to false (reset flag)
		log.Printf("Transition signal consumed and reset")
		// TODO return true

	}

	return false
}

// StartControlAPI starts the control API server on port 8080
func StartControlAPI() {

	// TODO: use http.HandleFunc() to create an endpoint at /switch, call handleSwitch

	log.Println("Starting Control API on :8080")
	go func() {

		// TODO: call http.ListenAndServe, pass :8080, and nil as args
		// TODO: perform error check for function call

	}()
}

func handleSwitch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: call TriggerTransition() on Manager (global flag)

	response := "Protocol transition triggered"

	json.NewEncoder(w).Encode(response)
}
