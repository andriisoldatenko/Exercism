// Package grains calculate the number of grains of wheat on a chessboard given that the number on each square doubles.
package grains

import "errors"

// Square returns number of grains of wheat on a chessboard for given number
func Square(n int) (sum uint64, err error) {
	sum = 1
	if n == 1 {
		return 1, nil
	}
	if n <= 0 {
		return 0, errors.New("square 0 returns an error")
	}
	if n > 64 {
		return 0, errors.New("square greater than 64 returns an error")
	}

	for i := 2; i <= n; i++ {
		sum = sum * 2
	}
	return sum, nil
}

// Total returns total of grains of wheat on a chessboard
func Total() (total uint64) {
	for i := 1; i <= 64; i++ {
		s, _ := Square(i)
		total += s
	}
	return
}
