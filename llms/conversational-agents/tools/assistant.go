package tools

import "context"

type Assistant struct{}

func (a Assistant) Name() string {
	return "AI Assistant"
}

func (a Assistant) Description() string {
	return `Useful for answering questions about the AI assistant
	Only questions about the name, age and creator of the AI can be answered. 

	An example of a question that fits this tool is: What is your name?

	Input should be name or age or creator depending on what was asked
	`
}

func (a Assistant) Call(ctx context.Context, input string) (string, error) {
	answers := map[string]string{
		"name":    "Gbenga AI",
		"age":     "1",
		"creator": "Gbenga Anifowoshe",
	}

	value, ok := answers[input]

	if ok {
		return value, nil
	}

	return "I cannot answer your question", nil
}
