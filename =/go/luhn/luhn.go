package luhn

import (
	"strconv"
	"strings"
)

func Valid(idNumber string) bool {

	idNumber = strings.ReplaceAll(idNumber, " ", "")

	if len(idNumber) <= 1 {
		return false
	}

	var err error

	var integerSum, intValue, doubledValue int

	for id := 0; id < len(idNumber); id++ {

		intValue, err = strconv.Atoi(string(idNumber[len(idNumber)-1-id]))

		if err != nil {
			return false
		}

		if id%2 == 0 {
			integerSum += intValue
			continue
		}

		doubledValue = intValue * 2

		if doubledValue > 9 {

			doubledValue = doubledValue - 9

		}

		integerSum += doubledValue

	}

	return integerSum%10 == 0

}
