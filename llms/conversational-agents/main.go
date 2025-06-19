package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	appTools "github.com/other-agents/tools"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/tools"
)

const modelName = "gemini-2.0-flash-exp"
const apiKey = "AIzaSyB-PzChELpgq2XfXzjg5XWYkhbmtlfkVkM"

var llmClient, _ = googleai.New(context.Background(), googleai.WithAPIKey(apiKey), googleai.WithDefaultModel(modelName))

func main() {
	agentTools := []tools.Tool{
		appTools.Assistant{},
		appTools.MyHotelLookup{},
	}

	// Conversational agents are agents that can use tools but are also very interactive since they have memory
	// These agents can reference its context to reply users. The context is always mantained in its memory
	agent := agents.NewConversationalAgent(llmClient, agentTools)

	memory := memory.NewConversationBuffer()
	executor := agents.NewExecutor(agent, agents.WithMemory(memory))

	ctx := context.Background()

	fmt.Printf("%s\n\n", "AI: Hey, I am GbengaAI assistant. Ask me anything")
	for {
		// Scan for the user input
		// Its better to use bufio as it allows us to wait for the user to press enter before reading. Scanlin potentially will read prompts and new lines
		fmt.Print("Question: ")
		os.Stdout.Sync() // This allows us to flush all FS in memory data to disk (NOT SUPER IMPORTANT HERE)

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// If the user types bye then exit the program
		if strings.ToLower(input) == "bye" {
			fmt.Println("AI: Bye for now 👋!")

			break
		}

		// Run the agent and output the answer of the Agent to StdOut
		output, err := chains.Run(ctx, executor, input)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("AI: %s\n\n", output)
	}
}
