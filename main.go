// Hello Forge is the simplest possible forge example.
//
// Shows how to call OpenAI using your OPENAI_API_KEY.
//
// Usage:
//
//	export OPENAI_API_KEY=sk-...
//	go run .
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	forge "github.com/katasec/forge-core"
	"github.com/katasec/forge-core/provider/openai"
	"github.com/katasec/forge-core/provider/xai"
)

func main() {

	// Setup config and context
	ctx := context.Background()
	agent := setupOpenAIAgent()
	//agent := setupXaiAgent()

	// Ask is the common path: user text in, AgentResponse out.
	resp, err := agent.Ask(ctx, "Hello! Who made you?")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.LastText())
	fmt.Printf("\n[tokens: %d in, %d out]\n", resp.Usage.InputTokens, resp.Usage.OutputTokens)
}

func setupOpenAIAgent() *forge.Agent {
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		log.Fatal("Set OPENAI_API_KEY environment variable")
	}

	// Setup config and context
	config := forge.Config{
		Provider:     openai.New(key, openai.ModelGPT54Nano),
		SystemPrompt: "You are a helpful assistant. Keep responses brief.",
	}

	// Create agent
	agent, err := forge.NewAgent(config)
	if err != nil {
		log.Fatal(err)
	}

	return agent
}

func setupXaiAgent() *forge.Agent {
	key := os.Getenv("XAI_API_KEY")
	if key == "" {
		log.Fatal("Set XAI_API_KEY environment variable")
	}

	// Setup config and context
	config := forge.Config{
		Provider:     xai.New(key, xai.ModelGrok4FastNonReasoning),
		SystemPrompt: "You are a helpful assistant. Keep responses brief.",
	}

	// Create agent
	agent, err := forge.NewAgent(config)
	if err != nil {
		log.Fatal(err)
	}

	return agent
}
