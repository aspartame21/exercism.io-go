package diffsquares

// Difference returns difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}

// SquareOfSum returns square of sum of first natural numbers
func SquareOfSum(num int) int {
	var res int

	for i := 1; i <= num; i++ {
		res += i
	}

	return res * res
}

// SumOfSquares returns sum of squares of first natural numbers
func SumOfSquares(num int) int {
	var res int

	for i := 1; i <= num; i++ {
		res += i * i
	}

	return res
}
