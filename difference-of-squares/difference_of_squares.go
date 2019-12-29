package diffsquares

// SquareOfSum calculate square of sum using formula as performance over iterating is uncomapable
func SquareOfSum(n int) int {
	s := (n * (n + 1) / 2)
	return s * s
}

// SumOfSquares calculate sum of squares using formula
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates diff of SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
