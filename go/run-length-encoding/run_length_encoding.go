// Package encode implmenets run lengts decode and encode
package encode

import (
	"fmt"
	"strconv"
	"unicode"
)

func RunLengthEncode(in string) string {
	runlengthEncoded := ""
	if len(in) == 0 {
		return runlengthEncoded
	}
	prev := ""
	count := 0
	for index, ch := range in {
		if prev == string(ch) || index == 0 {
			count++
		} else {
			if count >= 2 {
				runlengthEncoded = runlengthEncoded + fmt.Sprint(count) + prev
			} else {
				runlengthEncoded += prev
			}
			count = 1
		}
		prev = string(ch)

		if index == len(in)-1 {
			if count >= 2 {
				runlengthEncoded = runlengthEncoded + fmt.Sprint(count) + prev
			} else {
				runlengthEncoded += prev
			}
		}
	}
	return runlengthEncoded
}

func RunLengthDecode(in string) string {
	runlengthDecoded := ""
	if len(in) == 0 {
		return runlengthDecoded
	}
	digit := ""
	for _, ch := range in {
		if unicode.IsDigit(ch) {
			digit = digit + string(ch)
		} else {
			number, err := strconv.Atoi(digit)
			if err != nil {
				number = 1
			}
			for i := 0; i < number; i++ {
				runlengthDecoded += string(ch)
			}
			digit = ""
		}
	}
	return runlengthDecoded
}
