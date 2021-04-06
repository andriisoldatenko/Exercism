package raindrops

import "strconv"

// Convert converts a number into a string that contains raindrop sounds corresponding to certain potential factors.
func Convert(i int) string {
	result := ""
	if i%3 == 0 {
		result += "Pling"
	}
	if i%5 == 0 {
		result += "Plang"
	}
	if i%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(i)
	}
	return result
}
