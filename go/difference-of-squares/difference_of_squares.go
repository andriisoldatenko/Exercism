// Package diffsquares calculates diff squares
package diffsquares

// SquareOfSum returns square of sum starting from 1..n
func SquareOfSum(n int) int {
	// https://brilliant.org/wiki/sum-of-n-n2-or-n3/
	sum := (1 + n) * n / 2
	return sum * sum
}

// SumOfSquares returns sum of square from 1..n
func SumOfSquares(n int) int {
	// https://brilliant.org/wiki/sum-of-n-n2-or-n3/
	return (1 + n) * (2*n + 1) * n / 6
}

// Difference calculates diff between square of sum and sum of square
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
