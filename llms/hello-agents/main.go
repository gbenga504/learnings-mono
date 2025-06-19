package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/tools/wikipedia"
)

const modelName = "llama3.1:8b"

func run() error {
	ollamaClient, err := ollama.New(ollama.WithModel(modelName))

	if err != nil {
		return err
	}

	wikipedia := wikipedia.New("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")

	agentTools := []tools.Tool{
		tools.Calculator{},
		wikipedia,
	}

	agent := agents.NewOneShotAgent(ollamaClient, agentTools, agents.WithMaxIterations(3))

	executor := agents.NewExecutor(agent)

	question := "What is 2 multiplied by 0.23?"
	answer, err := chains.Run(context.Background(), executor, question)

	fmt.Println(answer)
	return err
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
