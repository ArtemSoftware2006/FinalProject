package main

import (
	"calculator/internal/agent"
)

func main() {
	agentCount := 5 // Number of agents to run
	for i := 0; i < agentCount; i++ {
		go agent.Run()
	}
	select {} // Block forever
}
