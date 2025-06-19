package customstring

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func computeSubstrings(currentSubstringIndexesHashMap map[string][]int, currentCharacter string, givenString string) map[string][]int {
	var result = make(map[string][]int)

	for charIndex, char := range givenString {
		character := string(char)

		// We skip doing anything if the current character does not match the with the lookup givenString char
		if currentCharacter != character {
			continue
		}

		// If the current substrings map is empty, then we just add the character to the result map
		if len(currentSubstringIndexesHashMap) == 0 {
			result[character] = append(result[character], charIndex)

			continue
		}

		// We go through the current substrings indexes hashmap and try to compute the new map
		for substring, substringIndexes := range currentSubstringIndexesHashMap {

			for _, substringIndex := range substringIndexes {
				// We save a new hashmap of the new substring and new indexes
				// The indexes is always the index of the substring in the given string
				if charIndex < substringIndex {
					newSubstring := givenString[charIndex : substringIndex+len(substring)]

					result[newSubstring] = append(result[newSubstring], charIndex)
				} else {
					newSubstring := givenString[substringIndex : charIndex+1]

					result[newSubstring] = append(result[newSubstring], substringIndex)
				}
			}
		}
	}

	return result
}

func calculateShortestSubstringWithAllChars(givenString string, characters []string) *string {
	// We track the substring and the last seen character index for each substring
	var substringIndexesHashMap = make(map[string][]int)

	for _, character := range characters {
		indexInGivenString := strings.Index(givenString, character)

		// If the character does not exist in the given string then return null since all the characters
		// must exist in the given string
		if indexInGivenString == -1 {
			return nil
		}

		// We compute the new substringIndexesHashMap
		substringIndexesHashMap = computeSubstrings(substringIndexesHashMap, character, givenString)
	}

	if len(substringIndexesHashMap) == 0 {
		return nil
	}

	// Check the shortest substring and return that instead
	var result string = ""

	for substring := range substringIndexesHashMap {
		if len(substring) < len(result) || len(result) == 0 {
			result = substring
		}
	}

	return &result
}

func GetShortestSubstringWithAllChars() {
	var givenString string
	var characters []string

	fmt.Println("=========== Program begins ==============")

	fmt.Println("Instructions")
	fmt.Println("=============")
	fmt.Println("Given a string and a set of characters, return the shortest substring containing all the characters in the set")

	fmt.Println()
	fmt.Println("Examples")
	fmt.Println("For example, given the string 'figehaeci' and the set of characters {a, e, i}, you should return 'aeci'")
	fmt.Println()
	fmt.Println("If there is no substring containing all the characters in the set, return null")
	fmt.Println()

	fmt.Println("Enter the string")
	fmt.Scan(&givenString)

	fmt.Println("Enter the characters")
	err := json.NewDecoder(os.Stdin).Decode(&characters)

	if err != nil {
		panic(err.Error())
	}

	result := calculateShortestSubstringWithAllChars(givenString, characters)

	fmt.Printf("The result is %v\n", result)
}
