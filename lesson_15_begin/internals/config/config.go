package config

import "time"


// AgentConfig holds all configuration values for the agent
type AgentConfig struct {
	ServerIP   string
	ServerPort string
	Timing     TimingConfig
	Protocol   string // this will be the starting protocol
	SharedSecret string // HMAC authentication key
}

// ServerConfig holds all configuration values for the server
type ServerConfig struct {
	ListeningInterface string
	ListeningPort      string
	Protocol           string // this will be the starting protocol
	TlsKey             string
	TlsCert            string
	SharedSecret string // HMAC authentication key
}

// TimingConfig holds timing-related configuration
type TimingConfig struct {
	Delay  time.Duration // Base delay between cycles
	Jitter int           // Jitter percentage (0-100)
}
