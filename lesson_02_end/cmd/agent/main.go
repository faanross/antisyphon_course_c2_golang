package main

import (
	"fmt"

	"c2framework/internals/agent"
	"c2framework/internals/config"
)

func main() {
	agentCfg := config.AgentConfig{
		Protocol: "https",
	}

	_, err := agent.NewAgent(&agentCfg)
	if err != nil {
		fmt.Println(err)
	}
}
