// Hello Forge is the simplest possible forge example.
//
// Shows how to call OpenAI, xAI, and Anthropic using Forge agents, including
// a tool call against each provider that supports tools.
//
// Usage:
//
//	export OPENAI_API_KEY=sk-...
//	export XAI_API_KEY=xai-...
//	export ANTHROPIC_API_KEY=sk-ant-...
//	go run .
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	forge "github.com/katasec/forge-core"
	"github.com/katasec/forge-core/provider/anthropic"
	"github.com/katasec/forge-core/provider/openai"
	"github.com/katasec/forge-core/provider/xai"
	"github.com/katasec/forge-core/tool"
)

// addInput is the typed argument struct for the add tool. Forge derives the
// tool's JSON Schema from it.
type addInput struct {
	A int `json:"a" jsonschema:"description=First number"`
	B int `json:"b" jsonschema:"description=Second number"`
}

// addTool returns a tool the model can call. It prints when it runs, so a real
// tool call is visible rather than inferred from the answer.
func addTool() forge.Tool {
	return tool.Func[addInput, int]("add", "Adds two numbers and returns their sum",
		func(_ context.Context, in addInput) (int, error) {
			fmt.Printf("  -> tool add(%d, %d) invoked\n", in.A, in.B)
			return in.A + in.B, nil
		})
}

func main() {
	ctx := context.Background()

	openAIAgent := setupOpenAIAgent()
	xaiAgent := setupXaiAgent()
	anthropicAgent := setupAnthropicAgent()

	runAgent(ctx, "OpenAI", "Hello! Who made you?", openAIAgent)
	runAgent(ctx, "xAI Search", "Find a recent source about xAI and summarize it briefly.", xaiAgent)
	runAgent(ctx, "Anthropic", "Hello! Who made you?", anthropicAgent)

	const toolPrompt = "What is 21 + 21? Use the add tool."
	runAgent(ctx, "OpenAI Tools", toolPrompt, openAIAgent)
	runAgent(ctx, "Anthropic Tools", toolPrompt, anthropicAgent)
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
		Tools:        []forge.Tool{addTool()},
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
		Provider:     xai.New(key, xai.ModelGrok4FastNonReasoning, xai.WithWebSearch()),
		SystemPrompt: "You are a helpful assistant. Keep responses brief.",
	}

	// Create agent
	agent, err := forge.NewAgent(config)
	if err != nil {
		log.Fatal(err)
	}

	return agent
}

func setupAnthropicAgent() *forge.Agent {
	key := os.Getenv("ANTHROPIC_API_KEY")
	if key == "" {
		log.Fatal("Set ANTHROPIC_API_KEY environment variable")
	}

	// Setup config and context
	config := forge.Config{
		Provider:     anthropic.New(key, anthropic.ModelClaudeOpus5),
		SystemPrompt: "You are a helpful assistant. Keep responses brief.",
		Tools:        []forge.Tool{addTool()},
	}

	// Create agent
	agent, err := forge.NewAgent(config)
	if err != nil {
		log.Fatal(err)
	}

	return agent
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
