// Package luhn defines luhn algorithms https://en.wikipedia.org/wiki/Luhn_algorithm
package luhn

import (
	"strings"
)

// Valid checks if input is valid luhn sequence
func Valid(in string) bool {
	trimed := strings.ReplaceAll(in, " ", "")
	length := len(trimed)
	if length <= 1 {
		return false
	}

	sum := 0
	counter := 0
	for i := length - 1; i >= 0; i-- {
		double := int(trimed[i] - '0')
		if double < 0 || double > 9 {
			return false
		}
		if counter%2 != 0 {
			double = 2 * double
			if double > 9 {
				double = double - 9
			}
		}
		sum += double
		counter++

	}
	return sum%10 == 0
}
