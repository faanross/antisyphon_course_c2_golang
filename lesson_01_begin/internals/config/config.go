package config

// AgentConfig holds all configuration values for the agent
type AgentConfig struct {

	// TODO: Add field ServerIP of type string
	// TODO: Add field ServerPort of type string
	// TODO: Add field Timing of type TimingConfig (i.e. another struct!)
	// TODO: Add field Protocol of type string
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
	// TODO: Add field Delay of type time.Duration
	// TODO: Add field Jitter of type int
	
}
