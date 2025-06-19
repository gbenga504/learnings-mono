package longestsubtring

import (
	"fmt"
	"slices"
)

func findLongestSubstring(sentence string, distinctCharacters int) string {
	if distinctCharacters == len(sentence) {
		return sentence
	}

	if distinctCharacters == 1 {
		return sentence[0:1]
	}

	var result string

	// We don't need to check if all the characters in the sentence can potentially
	// be the start of a substring, hence we take only useful characters
	// E.g if sentence = 'abc' and distinct characters is 2, then we will check only 'ab' as start characters
	sentenceToCheck := sentence[0 : len(sentence)-(distinctCharacters-1)]

	for index, startCharacter := range sentenceToCheck {
		// Keep track of the distinct characters for each iteration
		substring := string(startCharacter)
		var seenCharacters = []string{substring}

		// To find other characters that can be joined with the start character to form the substring
		// we add 1 to the index of the start character and check the full length of the sentence afterwards
		for _, otherChar := range sentence[index+1:] {
			// If the character has already been recorded, then we add it to the substring
			otherCharacter := string(otherChar)

			if slices.Index(seenCharacters, otherCharacter) != -1 {
				substring += otherCharacter
			}

			// We check if the other character is not in the distinct characters and adding it
			// will make sure that we don't surpass the number required for distinct characters, then we do
			// else we break from the forloop
			potentialSeenCharacters := append(seenCharacters, otherCharacter)

			if slices.Index(seenCharacters, otherCharacter) == -1 && len(potentialSeenCharacters) <= distinctCharacters {
				seenCharacters = potentialSeenCharacters
				substring += otherCharacter

			} else {
				break
			}
		}

		if len(substring) >= len(result) {
			result = substring
		}
	}

	return result
}

func Init() {
	var sentence string
	var distinctCharacters int

	fmt.Println("Program Begin")
	fmt.Println("=========================")
	fmt.Println("Enter the string")
	fmt.Scan(&sentence)

	fmt.Println("Enter the number of distinct characters")
	fmt.Scan(&distinctCharacters)

	result := findLongestSubstring(sentence, distinctCharacters)

	fmt.Printf("The result of the longest substring is %v\n", result)
}
