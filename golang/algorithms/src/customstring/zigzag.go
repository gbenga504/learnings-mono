package customstring

import (
	"fmt"
	"regexp"
)

func getzigzagForm(givenString string, numberOfLines int) string {
	lines := make([]string, numberOfLines)
	// We need to keep track of the distance X of each line characters
	distanceXOfLineChars := make([][]int, numberOfLines)
	result := ""

	currentLine := 0
	diagonalType := "right"

	for index, char := range givenString {
		character := string(char)

		// We get the distanceX of the current and previous character in the current line
		distanceXOfCurrentChar := index + 1
		distanceXOfPreviousChar := 0

		spaceRegexp := regexp.MustCompile(`\s`)
		numberOfCharsInCurrentLine := len(spaceRegexp.ReplaceAllString(lines[currentLine], ""))

		if numberOfCharsInCurrentLine == 0 {
			distanceXOfPreviousChar = 0
		} else {
			distanceXOfPreviousChar = distanceXOfLineChars[currentLine][numberOfCharsInCurrentLine-1]
		}

		numberOfSpaces := distanceXOfCurrentChar - distanceXOfPreviousChar

		// save the character to the current line
		// We will never have more than the total number of lines
		lines[currentLine] += fmt.Sprintf("%*s", numberOfSpaces, character)
		distanceXOfLineChars[currentLine] = append(distanceXOfLineChars[currentLine], distanceXOfCurrentChar)

		// We need to assign turning points at line 0 and if currentLine  == total number of lines - 1
		// We need to do this early enough so we know what direction diagonally we are moving in
		// For right diagonal, we wanna increase the current line and decrease if we are moving left
		if currentLine == 0 {
			diagonalType = "right"

		} else if currentLine == numberOfLines-1 {
			diagonalType = "left"

		}

		if diagonalType == "right" {
			currentLine++
		} else {
			currentLine--
		}

	}

	// fmt.Println(lines)
	for _, line := range lines {
		result += line + "\n"
	}

	return result
}

func ZigZag() {
	var givenString string
	var numberOfLines int

	fmt.Println("=========== Program begins ==============")

	fmt.Println("Instructions")
	fmt.Println("=============")
	fmt.Println("Given a string and a number of lines k, print the string in zigzag form. In zigzag, characters are printed out diagonally from top left to bottom right until reaching the kth line, then back up to top right, and so on")

	fmt.Println()
	fmt.Println("Examples")
	fmt.Println("For example, given the sentence 'thisisazigzag' and k = 4, you should print:")
	fmt.Println()
	fmt.Println(`
		t     a     g
		 h   s z   a
		  i i   i z
		   s     g
	`)
	fmt.Println()

	fmt.Println("Enter the string")
	fmt.Scan(&givenString)

	fmt.Println("Enter the number of lines")
	fmt.Scan(&numberOfLines)

	result := getzigzagForm(givenString, numberOfLines)

	fmt.Println()
	fmt.Print(result)
}
