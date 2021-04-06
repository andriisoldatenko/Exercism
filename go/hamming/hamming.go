package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("math: square root of negative number")
	}
	hammingDistance := 0
	for key, _ := range b {
		if a[key] != b[key] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
