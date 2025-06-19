package customstring

import (
	"fmt"
)

func getStartIndex(givenString string, pattern string) []int {
	var result []int

	for index := range givenString {
		if index > len(givenString)-len(pattern) {
			break
		}

		substring := givenString[index : index+len(pattern)]

		if substring == pattern {
			result = append(result, index)
		}
	}

	return result
}

func StartingIndicesOfOccurrence() {
	var givenString string
	var pattern string

	fmt.Println("=========== Program begins ==============")

	fmt.Println("Instructions")
	fmt.Println("=============")
	fmt.Println("Given a string and a pattern, find the starting indices of all occurrences of the pattern in the string")

	fmt.Println()
	fmt.Println("Examples")
	fmt.Println("For example, given the string 'abracadabra' and the pattern 'abr', you should return [0, 7]")
	fmt.Println()

	fmt.Println("Enter the string")
	fmt.Scan(&givenString)

	fmt.Println("Enter the pattern")
	fmt.Scan(&pattern)

	result := getStartIndex(givenString, pattern)

	fmt.Printf("The result is %v\n", result)
}
