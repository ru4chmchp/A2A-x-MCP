package main

import (
	"fmt"
	"log"

	"github.com/ru4chmchp/A2A-x-MCP/agentcard"
	"github.com/ru4chmchp/A2A-x-MCP/registry"
)

func main() {
	cards, err := agentcard.LoadAllAgentCards("registry/cards")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded %d agent cards\n", len(cards))
	agents := registry.NewRegistry()
	for _, card := range cards {

		if err := agentcard.ValidateAgentCard(card); err != nil {
			fmt.Println(" Invalid AgentCard:", err)
		} else {
			fmt.Println("Agent card valid")
		}
		agents.Register(card)

	}

	for _, result := range agents.FindByCategory("Reconnaissance") {
		println(result.Name)
	}

}
