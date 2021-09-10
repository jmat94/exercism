package diffsquares

// Function to return sum of all squared numbers
func SumOfSquares(number int) int {

	SumOfSquareValue := 0

	for num := 1; num <= number; num++ {

		SumOfSquareValue += (num * num)

	}

	return SumOfSquareValue

}

// Function to return the Square of the summed values
func SquareOfSum(number int) int {

	sumValue := 0

	for num := 1; num <= number; num++ {

		sumValue += num

	}

	return sumValue * sumValue

}

// Function to find the difference of the two
func Difference(number int) int {

	return SquareOfSum(number) - SumOfSquares(number)

}
