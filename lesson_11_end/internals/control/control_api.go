package control

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
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
	http.HandleFunc("/switch", handleSwitch)

	log.Println("Starting Control API on :8080")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Printf("Control API error: %v", err)
		}
	}()
}

func handleSwitch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	Manager.TriggerTransition()

	response := "Protocol transition triggered"

	json.NewEncoder(w).Encode(response)
}
