package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {

	// check if length of the string in 0 or empty
	if len(word) == 0 {
		return true
	}

	// convert string to lower case
	word = strings.ToLower(word)

	// create map to keep track of characters
	wordMap := make(map[rune]bool, len(word))

	for _, character := range word {

		if character == ' ' || character == '-' {
			continue
		}

		// check if character exists, if it does return false
		if wordMap[character] {
			return false
		}

		wordMap[character] = true
	}

	return true
}
