package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

// To use Langchain Go with llama models, we need an inference engine like Ollama
// Ollama is a gRPC server that wraps around the Llama models. Ollama does not need to be running though, it just needs to be installed and the model downloaded
func main() {
	// llm, err := ollama.New(ollama.WithModel("llama3.1:8b"))

	llm, err := googleai.New(context.Background(), googleai.WithAPIKey("AIzaSyB-PzChELpgq2XfXzjg5XWYkhbmtlfkVkM"), googleai.WithDefaultModel("gemini-2.0-flash-exp"))

	if err != nil {
		log.Fatal(err)
	}

	query := "Can you translate 'hello' to German"

	ctx := context.Background()

	fmt.Println("Final Response:")

	// For streaming the response, we need to pass a callback function
	_, err = llms.GenerateFromSinglePrompt(ctx, llm, query, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))

		return nil
	}))

	if err != nil {
		log.Fatal(err)
	}
}
