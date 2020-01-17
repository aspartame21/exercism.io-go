package diffsquares

// Difference returns difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}

// SquareOfSum returns square of sum of first natural numbers
func SquareOfSum(num int) int {
	sum := num * (num + 1) / 2
	return sum * sum
}

// SumOfSquares returns sum of squares of first natural numbers
func SumOfSquares(num int) int {
	return num * (num + 1) * (2*num + 1) / 6
}
