package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	llm, err := ollama.New(ollama.WithModel("llama3.1:8b"))

	if err != nil {
		log.Fatal(err)
	}

	texts := []string{
		"meteor",
		"comet",
		"puppy",
	}

	ctx := context.Background()
	embeddings, err := llm.CreateEmbedding(ctx, texts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got %d embeddings:\n", len(embeddings))

	for i, emb := range embeddings {
		fmt.Printf("%d: len=%d; first few=%v\n", i, len(emb), emb[:4])
	}
}
