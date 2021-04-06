package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	c := 0
	if n <= 0 {
		return 0, errors.New("zero is an error")
	}
	if n == 1 {
		return c, nil
	}
	for ; n > 1; c++ {
		if n&1 == 0 {
			n >>= 1
		} else {
			n += n<<1 + 1
		}
	}
	return c, nil
}
