package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {

	if number < 1 || number > 64 {
		return 0, errors.New("invalid number")
	}

	return uint64(math.Pow(2, float64(number-1))), nil

}

func Total() uint64 {

	var err error
	var total, currentValue uint64

	for i := 1; i <= 64; i++ {
		currentValue, err = Square(i)

		if err != nil {
			return 0
		}

		total += currentValue
	}

	return total

}
