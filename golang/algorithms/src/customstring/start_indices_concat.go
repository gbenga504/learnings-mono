package customstring

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

func getConcatIndex(givenString string, words []string) []int {
	var result []int
	wordLength := len(words[0])
	maxConcatLen := wordLength * len(words)
	lastIndex := len(givenString) - maxConcatLen

	for i := 0; i <= lastIndex; i++ {
		var seenWordIndex []int
		wordConcatStartIndex := i

		for {
			word := givenString[wordConcatStartIndex : wordConcatStartIndex+wordLength]

			wordIndex := slices.Index(words, word)

			// If the word does not exist in the set of words given then we want to exit the loop
			if wordIndex == -1 {
				break
			}

			// If the word index has already been seen, then we also exit the loop
			if slices.Index(seenWordIndex, wordIndex) != -1 {
				break
			}

			seenWordIndex = append(seenWordIndex, wordIndex)

			// If we have seen all the words in the string then we add [i] to the result and exit the loop
			if len(seenWordIndex) == len(words) {
				result = append(result, i)

				break
			}

			wordConcatStartIndex += wordLength
		}
	}

	return result
}

func GetStartingIndicesForConcat() {
	var givenString string
	var words []string

	fmt.Println("=========== Program begins ==============")

	fmt.Println("Instructions")
	fmt.Println("=============")
	fmt.Println("Given a string s and a list of words words, where each word is the same length, find all starting indices of substrings in s that is a concatenation of every word in words exactly once")

	fmt.Println()
	fmt.Println("Examples")
	fmt.Println("For example, given s = 'dogcatcatcodecatdog' and words = ['cat', 'dog'], return [0, 13], since 'dogcat' starts at index 0 and 'catdog' starts at index 13.")
	fmt.Println()
	fmt.Println("Given s = 'barfoobazbitbyte' and words = ['dog', 'cat'], return [] since there are no substrings composed of 'dog' and 'cat' in s. The order of the indices does not matter")
	fmt.Println()

	fmt.Println("Enter the string")
	fmt.Scan(&givenString)

	fmt.Println("Enter the words")
	err := json.NewDecoder(os.Stdin).Decode(&words)

	if err != nil {
		// We can also use log.Fatal
		fmt.Printf("An error occurred while reading the words ===> %v\n", err.Error())

		return
	}

	result := getConcatIndex(givenString, words)

	fmt.Printf("The result is %v\n", result)
}
