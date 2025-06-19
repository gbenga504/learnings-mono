package equalsentence

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

func isSynonyms(firstWord string, secondWord string, synonyms [][]string) bool {
	for _, synonym := range synonyms {
		if slices.Contains(synonym, firstWord) && slices.Contains(synonym, secondWord) {
			return true
		}
	}

	return false
}

func areSentenceEqual(firstSentence string, secondSentence string, synonyms [][]string) bool {
	firstSentenceArray := strings.Split(firstSentence, " ")
	secondSentenceArray := strings.Split(secondSentence, " ")

	if len(firstSentenceArray) != len(secondSentenceArray) {
		return false
	}

	var result bool = true

	for index, firstWord := range firstSentenceArray {
		secondWord := secondSentenceArray[index]

		if firstWord == secondWord {
			result = result && true
		} else {
			// Check if we have synonyms for the for these words
			result = result && isSynonyms(firstWord, secondWord, synonyms)
		}
	}

	return result
}

func Init() {
	var synonymsArray [][]string

	fmt.Println("Program Begin")
	fmt.Println("=========================")
	fmt.Println("Enter the synonyms array")

	// We cannot use jsonUnmarshal when reading from stdIn because the value from stdIn isn't a properly json encoded
	// json.NewDecoder works perfectly
	jsonErr := json.NewDecoder(os.Stdin).Decode(&synonymsArray)

	if jsonErr != nil {
		fmt.Printf("An error occurred while passing JSON %v", jsonErr.Error())

		return
	}

	result := areSentenceEqual("He wants to eat food.", "He wants to consume food.", synonymsArray)

	fmt.Printf("The result of if sentence are equal is %v\n", result)
}
