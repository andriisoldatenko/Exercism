// Package isogram calculate if string isogram or not
package isogram

import (
	"unicode"
)

// IsIsogram returns true if input string has unique list of chars except hyphens and spaces
func IsIsogram(input string) bool {
	uniqueChars := make(map[rune]bool)
	for _, ch := range input {
		if ch == '-' || ch == ' ' {
			continue
		}
		lowerCh := unicode.ToLower(ch)
		if uniqueChars[lowerCh] {
			return false
		}
		uniqueChars[lowerCh] = true

	}
	return true
}
