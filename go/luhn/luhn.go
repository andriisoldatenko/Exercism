// Package luhn defines luhn algorithms https://en.wikipedia.org/wiki/Luhn_algorithm
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// SpaceStringsBuilder strips all whitespace from a string (performs a single allocation, but may grossly overallocate if the source string is mainly whitespace.)
func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

// Valid checks if input is valid luhn sequence
func Valid(in string) bool {
	numbers := []int{}
	trimed := SpaceStringsBuilder(in)
	if len(trimed) <= 1 {
		return false
	}

	for _, ch := range trimed {
		if num, err := strconv.Atoi(string(ch)); err != nil {
			return false
		} else {
			numbers = append([]int{num}, numbers...)
		}
	}
	sum := 0
	for i := len(numbers) - 1; i >= 0; i -= 1 {
		double := numbers[i]
		if i%2 != 0 {
			double = 2 * double
			if double > 9 {
				double = double - 9
			}
		}
		sum += double

	}
	if sum%10 != 0 {
		return false
	}
	return true
}
