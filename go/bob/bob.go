// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"strings"
)

var (
	asciiLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	switch {
	case strings.HasSuffix(remark, "?") && strings.ToUpper(remark) == remark && strings.ContainsAny(remark, asciiLetters):
		return fmt.Sprint("Calm down, I know what I'm doing!")
	case strings.HasSuffix(strings.TrimSpace(remark), "?"):
		return fmt.Sprint("Sure.")
	case strings.ToUpper(remark) == remark && strings.ContainsAny(remark, asciiLetters):
		return fmt.Sprint("Whoa, chill out!")
	case !strings.ContainsAny(remark, asciiLetters+digits):
		return fmt.Sprint("Fine. Be that way!")
	default:
		return fmt.Sprint("Whatever.")
	}
}
