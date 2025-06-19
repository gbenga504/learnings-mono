package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google-gen-ai/db"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		log.Fatalln("Environment variable GEMINI_API_KEY not set")
	}

	//////////////////THIS IS FAST AS FAR AS THE CODE HAS BEEN BUILT////////////////////////
	fmt.Println("Loaded environment variable")
	//////////////////////////////////////////

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	var as map[string]any = map[string]any{"name": "asas"}

	///////////////THIS IS FAST///////////////////////////
	fmt.Println("Created genimi client")
	//////////////////////////////////////////

	model := client.GenerativeModel("gemini-2.0-flash")

	//////////// GET FLIGHTS TOOL ////////////////////////
	// A tool is used to group similar set of function together
	// This is useful when we want to use the model to call functions
	FlightsFunction := &genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{
			{
				Name:        "getAvailableFlights",
				Description: "Get available flights for a given route and date. This also returns the price of each flight so you can book the cheapest one",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"origin":      {Type: genai.TypeString, Description: "Departure city"},
						"destination": {Type: genai.TypeString, Description: "Arrival city"},
						"date":        {Type: genai.TypeString, Description: "Flight date (YYYY-MM-DD)"},
					},
					Required: []string{"origin", "destination", "date"},
				},
			},
			{
				Name:        "bookFlight",
				Description: "Book a specific flight using the flightId. This returns the bookingId and status of the booking",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"flightId": {Type: genai.TypeString, Description: "Unique identifier for the flight"},
					},
					Required: []string{"flightId"},
				},
			},
		},
	}

	// The tools allows the model use the tools explicitly rather than rely on its own knowledge
	model.Tools = []*genai.Tool{FlightsFunction}

	// Ranges from 0 - 2 in most cases. The lower the temperature, the more conservative the model is.
	// Temperature is only applied after setToP and setToK.
	// The higher the temperature, the more creative the model is.
	// Bteween 0 - 2
	model.SetTemperature(1)
	// Sets the maximum number of probable tokens that can be considred in the output.
	// The higher the value, the more tokens the model can consider in the output.
	// The lower the value, the less tokens the model can consider in the output.
	// Set to a lower value if you want a more conservative output.
	// Set to a higher value if you want a more creative output.
	model.SetTopK(40)
	// Sets the probability of the model considering the top P tokens in the output.
	// This adds the probability of each token from (setToK) until the prbability is slighly equal to setToP.
	// The higher the value, the more creative the output.
	// The lower the value, the more conservative the output.
	model.SetTopP(0.95)
	// Sets the maximum number of tokens in the output. The higher the output tokens, the longer the response could be
	// The lower the value, the shorter the response could be.
	model.SetMaxOutputTokens(8192)
	// We want this to be text/plain for chat apps
	model.ResponseMIMEType = "text/plain"
	// The CandidateCount is the number of candidates that the model will generate.
	// These candidates are just responses that tge model thinks are the best.
	// However, it is pointless ATM because the model only generates one candidate.
	// model.GenerationConfig.CandidateCount = 1

	// System instructions are always super maga helpful in directing the model on what to do
	// Always provide system instructions to the model
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(`
			You are a flight assistant. Your role is to help find flights for users and book them if asked.

			If you are asked to book a flight, always ask for confirmation with details of the flight before confirmating it.
			Also be very nice to users, use emojis when necessary but don't over use them and make sure to provide the best flight options.

			If you are asked to make suggestions about places one can spend a vacation, make sure you try to understand what the user enjoys, since this will help you make better suggestions
		`)},
	}

	session := model.StartChat()
	session.History = db.TransformToGenAIContents() // Get the chat history from the DB

	///// INITIAL PROMPT /////
	prompt := genai.Text("Can you help book me the cheapest flight to London for the 24th December 2025 ? I currently stay in Berlin. Also just book the flight, no need to get confirmation from me")

	////// Save user's prompt in the DB /////
	userHistory := db.TransformToChatHistory(&genai.Content{
		Role:  "user",
		Parts: []genai.Part{prompt},
	})
	db.SaveChatHistory(userHistory)

	messageStream := session.SendMessageStream(ctx, prompt)
	resp, err := messageStream.Next()

	for resp != nil {
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}

		/////// STREAM TEXT RESPONSE BACK TO USERS /////////
		// We check if the model generated some textual response
		// If it did, we want to print it out so the user is aware of what is happening
		// Here we iterate through all the parts of the response and print only text parts
		for _, part := range resp.Candidates[0].Content.Parts {
			if text, ok := part.(genai.Text); ok {
				if strings.TrimSpace(string(text)) == "" {
					continue
				}

				fmt.Printf("%s", text)
			}
		}

		resp, err = messageStream.Next()
		mergedResp := messageStream.MergedResponse()

		////// STREAM IS DONE AND WE NEED TO CALL FUNCTIONS IF ANY /////////
		if resp == nil && len(mergedResp.Candidates[0].FunctionCalls()) > 0 {
			/////// Save ai's prompt in the DB ///////
			aiHistory := db.TransformToChatHistory(&genai.Content{
				Role:  "ai",
				Parts: mergedResp.Candidates[0].Content.Parts,
			})
			db.SaveChatHistory(aiHistory)

			var toolResponse = make(map[string]any)

			for _, fc := range mergedResp.Candidates[0].FunctionCalls() {
				// In production code, we want to run a go routine to handle the function calls,
				// and save the responses in toolResponse (concurrency)
				switch fc.Name {
				case "getAvailableFlights":
					toolResponse["getAvailableFlights"] = simulateGetAvailableFlights()
				case "bookFlight":
					toolResponse["bookFlight"] = simulateBookFlight()
				}
			}

			////// SAVE CHAT HISTORY AND SEND TOOL RESPONSE TO MODEL ///////////
			parts := []genai.Part{prompt}
			parts = append(parts, toolResponseToFunctionResponse(toolResponse)...)

			systemHistory := db.TransformToChatHistory(&genai.Content{
				Role:  "system",
				Parts: parts,
			})
			db.SaveChatHistory(systemHistory)

			messageStream = session.SendMessageStream(ctx, parts...)
			resp, err = messageStream.Next()
		}
	}
}

func LogPart(part genai.Part) {
	switch v := part.(type) {
	case genai.FunctionCall:
		fmt.Printf("FUNCTION CALL ===> %+v\n", v)
	case genai.Text:
		fmt.Printf("TEXT CALL ===> %+v\n", v)
	default:
		fmt.Printf("OTHERS ===> %+v\n", v)
	}
}

func simulateGetAvailableFlights() []byte {
	resp, err := json.Marshal([]map[string]any{
		{"flightId": "F1", "price": 500, "airline": "Airline A"},
		{"flightId": "F2", "price": 450, "airline": "Airline B"},
		{"flightId": "F3", "price": 550, "airline": "Airline C"},
	},
	)

	if err != nil {
		log.Fatalf("Error marshalling response: %v", err)
	}

	return resp
}

func simulateBookFlight() []byte {
	resp, err := json.Marshal(map[string]any{
		"bookingId": "B12345",
		"status":    "confirmed",
		"flightId":  "F2",
	})

	if err != nil {
		log.Fatalf("Error marshalling response: %v", err)
	}

	return resp
}

func toolResponseToFunctionResponse(toolResponse map[string]any) []genai.Part {
	var result []genai.Part

	for toolName, response := range toolResponse {
		result = append(result, &genai.FunctionResponse{
			Name:     toolName,
			Response: map[string]any{"response": response},
		})
	}

	return result
}
