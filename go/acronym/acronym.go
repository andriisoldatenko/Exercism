// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[a-zA-Z']+`)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	words := re.FindAllString(s, -1)
	acronym := ""
	for _, word := range words {
		acronym += strings.ToUpper(string(word[0]))
	}
	return acronym
}
