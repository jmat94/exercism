package scrabble

import "strings"

// Function to calculate the score of the string given as an input
func Score(word string) int {

	// a map to store alphabets along with their values
	var alphabetValues = map[byte]int{

		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
		'L': 1,
		'N': 1,
		'R': 1,
		'S': 1,
		'T': 1,
		'D': 2,
		'G': 2,
		'B': 3,
		'C': 3,
		'M': 3,
		'P': 3,
		'F': 4,
		'H': 4,
		'V': 4,
		'W': 4,
		'Y': 4,
		'K': 5,
		'J': 8,
		'X': 8,
		'Q': 10,
		'Z': 10,
	}

	// variable to keep track of the value of each character in the string
	score := 0

	// converting the string to Uppercase
	upperCaseWord := strings.ToUpper(word)

	for i := range upperCaseWord {
		score += alphabetValues[upperCaseWord[i]] // Updating the score for each character seen
	}

	return score

}
