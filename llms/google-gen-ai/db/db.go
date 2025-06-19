package db

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/google/generative-ai-go/genai"
)

// //////////// DB TYPES ///////////////
// These fields are not exported but we need to be able to marshal and unmarschal them to JSON
// So we use the `json` tag to specify the field name when marshalling and unmarshalling
// Fields need to be exported before they can be marschalled and unmarschal
type data struct {
	Action string `json:"action"` // this will be text | functionCall | functionResponse. Use enums in production code
	// this will be the text or function call or function response
	// For function call, we will have the function name and args
	// For function response, we will have the entire response as a map[string]any
	Content string `json:"content"`
}

type message struct {
	Role         string `json:"role"` // this will be either "user" | "system" | "ai". Use enums in production code
	Data         []data `json:"data"`
	ModelVersion string `json:"modelVersion"`
}

type chatHistory struct {
	// We
	// In production, we will have fields like userID, createdAt etc
	// map[string]any{ID, UserID, SessionID, CreatedAt, UpdatedAt}
	Id        string  `json:"id"` // DB _id
	SessionId string  `json:"sessionId"`
	Message   message `json:"message"`
}

/////////// END OF DATABASE TYPES ////////////////

// This will live outside the repository and will handle transforming the chat history into a format that can be saved in the DB
// this will be called to convert genai Content to DB chatHistory before we save in the DB
func TransformToChatHistory(content *genai.Content) chatHistory {
	newChatHistory := chatHistory{
		Id:        strconv.Itoa(rand.Int()), // Mongo db _id
		SessionId: "gbenga",                 // A proper session id will be generated and sent by the frontend
		Message: message{
			ModelVersion: "gemini-2.0-flash", // This will be passed to this function
			Data:         []data{},
		},
	}

	////// Set the role ///////
	switch content.Role {
	case "model":
		newChatHistory.Message.Role = "ai"
	case "system":
		newChatHistory.Message.Role = "system"
	default:
		newChatHistory.Message.Role = "user"
	}

	/////// Set the data //////
	var messageData []data

	for _, part := range content.Parts {
		// Maybe a switch is better but "if" gives us better type ascertion
		if text, ok := part.(genai.Text); ok {
			messageData = append(messageData, data{Action: "text", Content: string(text)})
		}

		// In production code, we want to marschal structs instead of map[string]any for better types
		// So we can unmarschal also well
		if funcCall, ok := part.(genai.FunctionCall); ok {
			stringifiedData, err := json.Marshal(map[string]any{
				"name": funcCall.Name,
				"args": funcCall.Args,
			})

			if err != nil {
				log.Fatalf("An error occurred when parsing json %s\n", err.Error())
			}

			messageData = append(messageData, data{Action: "functionCall", Content: string(stringifiedData)})
		}

		if funcResp, ok := part.(genai.FunctionResponse); ok {
			stringifiedData, err := json.Marshal(map[string]any{
				"name":     funcResp.Name,
				"response": funcResp.Response,
			})

			if err != nil {
				log.Fatalf("An error occurred when parsing json %s\n", err.Error())
			}

			messageData = append(messageData, data{Action: "functionResponse", Content: string(stringifiedData)})
		}
	}

	newChatHistory.Message.Data = messageData

	return newChatHistory
}

// This will read all chat histories from DB and convert to an array of genai Contents
func TransformToGenAIContents() []*genai.Content {
	chatHistories := readChatHistoriesFromDB()

	var result []*genai.Content

	for _, ch := range chatHistories {
		role := "model"

		if ch.Message.Role == "system" {
			role = "user"
		}

		var parts []genai.Part
		for _, d := range ch.Message.Data {
			if d.Action == "text" {
				parts = append(parts, genai.Text(d.Content))
			} else if d.Action == "functionCall" {
				var funcCall struct {
					Name string
					Args map[string]any
				}

				if err := json.Unmarshal([]byte(d.Content), &funcCall); err != nil {
					log.Fatalf("Cannot unmarschal funcCall d %s", err.Error())
				}

				parts = append(parts, genai.FunctionCall{
					Name: funcCall.Name,
					Args: funcCall.Args,
				})
			} else if d.Action == "functionResponse" {
				var funcResp struct {
					Name     string
					Response map[string]any
				}

				if err := json.Unmarshal([]byte(d.Content), &funcResp); err != nil {
					log.Fatalf("Cannot unmarschal funcResp d %s", err.Error())
				}

				parts = append(parts, genai.FunctionResponse{
					Name:     funcResp.Name,
					Response: funcResp.Response,
				})
			}
		}

		content := genai.Content{
			Role:  role,
			Parts: parts,
		}

		result = append(result, &content)
	}

	return result
}

func SaveChatHistory(ch chatHistory) {
	chatHistories := readChatHistoriesFromDB()

	chatHistories = append(chatHistories, ch)

	jm, err := json.Marshal(chatHistories)

	if err != nil {
		log.Fatalf("Cannot marschal chat history when saving chat %s", err.Error())
	}

	if err := os.WriteFile("./db/chatHistory.json", jm, 0644); err != nil {
		log.Fatalf("Cannot write to database %s", err.Error())
	}
}

func readChatHistoriesFromDB() []chatHistory {
	// Go sets the working directory to the path thatg the "go run" command is run from
	// hence, this file needs to be relative based on the path of the working directory
	fileContent, err := os.ReadFile("./db/chatHistory.json")

	if err != nil {
		log.Fatalf("Error reading file: %v", err.Error())
	}

	var chatHistories []chatHistory
	if err = json.Unmarshal(fileContent, &chatHistories); err != nil {
		log.Fatalf("Cannot read chat histories from database %s\n", err.Error())
	}

	return chatHistories
}
