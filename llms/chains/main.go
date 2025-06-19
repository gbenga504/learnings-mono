package main

import (
	"context"
	"fmt"
	"log"

	appTools "github.com/chains/tools"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/tools"
)

func main() {
	ctx := context.Background()
	llmClient, _ := googleai.New(ctx, googleai.WithAPIKey("AIzaSyB-PzChELpgq2XfXzjg5XWYkhbmtlfkVkM"), googleai.WithDefaultModel("gemini-2.0-flash-exp"))

	a := agents.NewOpenAIFunctionsAgent()

	//
	// MATH CHAIN
	// This LLM has an inbuilt prompt for math based calculation and uses the starlark calculator to evaluate expression
	// This chain is super good for maths but bad for everything else. Hence do not use for anything else
	//
	output, err := chains.Run(ctx, chains.NewLLMMathChain(llmClient), "What is 3 + 6")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)

	//
	// LLM Chain
	// This is a simple chain with an LLM plugin
	//
	prompt := prompts.NewPromptTemplate(`
		You are an AI assistant.

		Do not spend a lot of time trying to find the answer, if you don't know the answer, just say "Sorry, I don't have an answer"

		Question: {{.question}}
	`, []string{"question"})

	chains.Run(ctx, chains.NewLLMChain(llmClient, prompt), "Who is Aliko Dangote?", chains.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))

		return nil
	}))

	//
	// Simple Sequential Chain
	//
	// This is a chain that runs multiple chains in sequence, the output of a chain is an input to another
	// Thus all thechains must have 1 input and 1 output
	// We will add 2 agents to our sequential chain.
	// 1 Agent generates a random number and the other other tells us if the number generated is a positive or negative numberp
	randomNumberAgent := agents.NewOneShotAgent(llmClient, []tools.Tool{appTools.NumRandomizerTool{}})
	signAgent := agents.NewOneShotAgent(llmClient, []tools.Tool{appTools.SignTool{}})

	randomNumberAgentExecutor := agents.NewExecutor(randomNumberAgent)
	signAgentExecutor := agents.NewExecutor(signAgent)

	newSimpleSequentialChain, _ := chains.NewSimpleSequentialChain([]chains.Chain{randomNumberAgentExecutor, signAgentExecutor})

	output, err = chains.Run(ctx, newSimpleSequentialChain, "Can you generate a random number?")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)

	//
	// Sequential Chain
	//
	// For more complex sequential chain, it is best to define your prompts yourself with the sub chains here
	// so you can control your inputs and outputs
	// Hence this example with not work properly at the moment
	//
	randomNumberAgent.OutputKey = "input"
	signAgent.OutputKey = "result"
	newSequentialChain, err := chains.NewSequentialChain([]chains.Chain{randomNumberAgentExecutor, signAgentExecutor}, []string{"input"}, []string{"result"})

	if err != nil {
		log.Fatal(err)
	}

	output, err = chains.Run(ctx, newSequentialChain, "Can you generate a random number?")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
