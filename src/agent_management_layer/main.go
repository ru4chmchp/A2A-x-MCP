package main

import (
	"fmt"
	"log"

	"github.com/ru4chmchp/A2A-x-MCP/agentcard"
)

func main() {
	cards, err := agentcard.LoadAllAgentCards("registry/cards")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded %d agent cards\n", len(cards))

	for _, card := range cards {

		if err := agentcard.ValidateAgentCard(card); err != nil {
			fmt.Println(" Invalid AgentCard:", err)
		} else {
			fmt.Println("Agent card valid")
		}

	}

}
