package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Langchain go has an OpenAI functions that this is only compatible with the OPen AI model
// Hence we use access the functions provided by gemini using their go lib
// We are building a light control system using Gemini Function calling
// Function calling provided by AI models is more faster, optimized, intelligent and uses less tokens than using a OneShotZeroAgent
// or a ConversationalAgent with tools. Hence when available we should use them especially because of the advantages they provide
func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyB-PzChELpgq2XfXzjg5XWYkhbmtlfkVkM"))

	if err != nil {
		log.Fatal(err)
	}

	// Select a Gemini model here
	model := client.GenerativeModel("gemini-2.0-flash-exp")

	// Create the Function declaration that can be passed to the gen model as a tool
	lightControlTool := &genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{{
			// This should always be concise. We should never use spaces etc.
			// Use camelCase or underscore to seperate multiple words
			Name: "controlLight",
			// This should also be super clear
			Description: "Set the brightness and color temperature of a room light.",
			Parameters: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"brightness": {
						Type:        genai.TypeString,
						Description: "Light level from 0 to 100. Zero is off and" + " 100 is full brightness.",
					},
					"colorTemperature": {
						Type:        genai.TypeString,
						Description: "Color temperature of the light fixture which" + " can be `daylight`, `cool` or `warm`.",
					},
				},
				Required: []string{"brightness", "colorTemperature"},
			},
		}},
	}

	// Add the light control to the model tools
	model.Tools = []*genai.Tool{lightControlTool}

	/////////
	//////// HERE WE START THE PROCESSING
	///////

	// start a new chat session
	session := model.StartChat()

	prompt := "Dim the lights so the room feels cozy and warm."

	// Send the message to the generative model
	resp, err := session.SendMessage(ctx, genai.Text(prompt))

	if err != nil {
		log.Fatal(err)
	}

	// Check that you got the expected function call back
	part := resp.Candidates[0].Content.Parts[0]
	funcall, ok := part.(genai.FunctionCall)

	if !ok {
		log.Fatalf("Expected type FunctionCall, got %T", part)
	}

	if g, e := funcall.Name, lightControlTool.FunctionDeclarations[0].Name; g != e {
		log.Fatalf("Expected FunctionCall.Name %q, got %q", e, g)
	}

	fmt.Printf("Received function call response:\n%q\n\n", part)

	brightness, _ := funcall.Args["brightness"].(string)
	colorTemp, _ := funcall.Args["colorTemperature"].(string)

	bright, _ := strconv.Atoi(brightness)

	// Send the hypothetical API result back to the generative model
	apiResult := setLightValues(bright, colorTemp)

	resp, err = session.SendMessage(ctx, genai.FunctionResponse{
		Name:     lightControlTool.FunctionDeclarations[0].Name,
		Response: apiResult,
	})

	if err != nil {
		log.Fatalf("Error sending message: %v\n", err)
	}

	// Show the model's response, which is expected to the text.
	// for _, part := range resp.Candidates[0].Content.Parts {
	// 	fmt.Printf("%v\n", part)
	// }
}

func setLightValues(brightness int, colorTemp string) map[string]any {
	// This is a mock API that returns the requested lightening values
	return map[string]any{
		"brightness": brightness,
		"colorTemp":  colorTemp,
	}
}
