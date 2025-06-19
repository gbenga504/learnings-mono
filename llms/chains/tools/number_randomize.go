package tools

import (
	"context"
	"math/rand"
	"strconv"
)

type NumRandomizerTool struct{}

func (n NumRandomizerTool) Name() string {
	return "numberRandomizer"
}

func (n NumRandomizerTool) Description() string {
	return `Useful for generating random numbers`
}

func (n NumRandomizerTool) Call(ctx context.Context, input string) (string, error) {
	minNumber := -10
	maxNumber := 10

	randomNum := rand.Intn(maxNumber-minNumber+1) + minNumber

	return strconv.Itoa(randomNum), nil
}
