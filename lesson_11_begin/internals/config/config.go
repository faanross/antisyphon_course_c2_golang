package config

import "time"

// SharedSecret is the key used for HMAC authentication
// In production, use a cryptographically random key and never commit to source control
const SharedSecret = "your-super-secret-key-change-in-production"

// AgentConfig holds all configuration values for the agent
type AgentConfig struct {
	ServerIP   string
	ServerPort string
	Timing     TimingConfig
	Protocol   string // this will be the starting protocol
}

// ServerConfig holds all configuration values for the server
type ServerConfig struct {
	ListeningInterface string
	ListeningPort      string
	Protocol           string // this will be the starting protocol
	TlsKey             string
	TlsCert            string
}

// TimingConfig holds timing-related configuration
type TimingConfig struct {
	Delay  time.Duration // Base delay between cycles
	Jitter int           // Jitter percentage (0-100)
}
