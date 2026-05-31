// Hello Forge is the simplest possible forge example.
//
// Shows how to call OpenAI and xAI using Forge agents.
//
// Usage:
//
//	export OPENAI_API_KEY=sk-...
//	export XAI_API_KEY=xai-...
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
	ctx := context.Background()

	runAgent(ctx, "OpenAI", "Hello! Who made you?", setupOpenAIAgent())
	runAgent(ctx, "xAI Search", "Find a recent source about xAI and summarize it briefly.", setupXaiAgent())
}

func runAgent(ctx context.Context, name string, prompt string, agent *forge.Agent) {
	resp, err := agent.Ask(ctx, prompt)
	if err != nil {
		log.Fatal(err)
	}

	printResponse(name, resp)
}

func printResponse(name string, resp *forge.AgentResponse) {
	fmt.Printf("\n[%s]\n", name)
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
		Provider: xai.New(
			key,
			xai.ModelGrok4FastNonReasoning,
			xai.WithWebSearch(),
		),
		SystemPrompt: "You are a helpful assistant. Keep responses brief.",
	}

	// Create agent
	agent, err := forge.NewAgent(config)
	if err != nil {
		log.Fatal(err)
	}

	return agent
}
