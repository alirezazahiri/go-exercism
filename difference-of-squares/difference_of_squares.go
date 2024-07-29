package diffsquares

import "math"

func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := n * (n + 1) * (2*n + 1) / 6
	return sum
}

func Difference(n int) int {
	return int(math.Abs(float64(SumOfSquares(n)) - float64(SquareOfSum(n))))
}
