package tools

import (
	"context"
	"strconv"
)

type Sign string

const (
	Positive Sign = "Positive"
	Negative Sign = "Negative"
)

type SignTool struct{}

func (n SignTool) Name() string {
	return "checkSign"
}

func (n SignTool) Description() string {
	return `Useful for checking the sign of a number.
	This tool tells if a number is a positive or negative.

	E.g
	Given the number 9, the tool returns Positive

	---
	The input to this tool is the number to check
	`
}

func (n SignTool) Call(ctx context.Context, input string) (string, error) {
	number, _ := strconv.Atoi(input)

	if number >= 0 {
		return string(Positive), nil
	}

	return string(Negative), nil
}
