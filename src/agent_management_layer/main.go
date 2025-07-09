package main

// "fmt"
// "log"

// "github.com/ru4chmchp/A2A-x-MCP/agentcard"
// "github.com/ru4chmchp/A2A-x-MCP/registry"
import (
	"github.com/ru4chmchp/A2A-x-MCP/task"
)

func main() {
	// cards, err := agentcard.LoadAllAgentCards("registry/cards")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Loaded %d agent cards\n", len(cards))
	// agents := registry.NewRegistry()
	// for _, card := range cards {

	// 	if err := agentcard.ValidateAgentCard(card); err != nil {
	// 		fmt.Println(" Invalid AgentCard:", err)
	// 	} else {
	// 		fmt.Println("Agent card valid")
	// 	}
	// 	agents.Register(card)

	// }

	// for _, result := range agents.FindByCategory("Reconnaissance") {
	// 	println(result.Name)
	// }

	store := &task.TaskStorage{FilePath: "task/tasks.json"}
	loadedTasks, err := store.Load()

	if err != nil {
		println("")
	}

	manager := task.NewManager()

	manager.ImportTasks(loadedTasks)

	for _, task := range manager.ListTasks() {
		println(task.AgentID)
	}

}
