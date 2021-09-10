package raindrops

import (
	"strconv"
)

func Convert(number int) string {

	// create an empty string to store the result
	result := ""

	// check if the number is a factor of 3
	if number%3 == 0 {
		result += "Pling"
	}

	// check if the number is a factor of 5
	if number%5 == 0 {
		result += "Plang"
	}

	// check if the number is a factor of 7
	if number%7 == 0 {
		result += "Plong"
	}

	// return the number if it is not divisible by 3 or 5 or 7
	if result == "" {
		return strconv.Itoa(number)
	}

	return result

}
