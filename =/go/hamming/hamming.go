package hamming

import "errors"

// function to calculate the hamming distance
func Distance(a, b string) (int, error) {

	// checking if lengths of the two strings are the same
	if len(a) == len(b) {
		errorCount := 0

		// iterating through the first string
		for index := range a {

			// if character in the first string is not equal to the second string at the same location increase
			// the errorCount by one
			if a[index] != b[index] {
				errorCount += 1
			}

		}
		return errorCount, nil
	}

	return 0, errors.New("incompatible lengths")

}
