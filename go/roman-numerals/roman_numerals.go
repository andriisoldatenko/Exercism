package romannumerals

import (
	"errors"
)

var baseNums = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func ToRomanNumeral(number int) (string, error) {
	if number <= 0 || number > 3000 {
		return "", errors.New("error: invalid roman number")
	}
	nums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	syms := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	roman := ""
	i := 12
	for number > 0 {
		div := number / nums[i]
		number = number % nums[i]
		for div > 0 {
			roman += syms[i]
			div--
		}
		i--
	}

	return roman, nil
}
