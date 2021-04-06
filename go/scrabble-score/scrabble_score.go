package scrabble

import "strings"

// Letter                           Value
// A, E, I, O, U, L, N, R, S, T       1
// D, G                               2
// B, C, M, P                         3
// F, H, V, W, Y                      4
// K                                  5
// J, X                               8
// Q, Z                               10

func Score(word string) int {
	runes := []rune(strings.ToUpper(word))
	score := 0
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}
	return score
}
