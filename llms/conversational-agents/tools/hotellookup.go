package tools

import "context"

type MyHotelLookup struct{}

func (h MyHotelLookup) Name() string {
	return "Hotel History Lookup"
}

func (h MyHotelLookup) Description() string {
	return `Useful for looking up or searching for the history of an hotel
	
	Examples of questions that fits this tool are:
	1. Can you tell me about the Sheraton Hotel ?
	2. What do you know about Ibiza hotel ?

	Input should be the name of the hotel converted to lowercase e.g "sheraton"
	If you don't have sufficient information to use this tool e.g the user does not provide the name of the hotel, ask them for it
	`
}

func (h MyHotelLookup) Call(ctx context.Context, input string) (string, error) {
	history := map[string]string{
		"sheraton": "This is the largest hotel in Lagos",
		"ibiza":    "This is the largest hotel in spain",
	}

	value, ok := history[input]

	if ok {
		return value, nil
	}

	return "Cannot tell you anything about this hotel", nil
}
