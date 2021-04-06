// Package summultiples
package summultiples

// SumMultiples
func SumMultiples(limit int, divisors ...int) int {
	nums := make(map[int]bool)
	sum := 0
	for _, d := range divisors {
		if d != 0 {
			for k := d; k < limit; k += d {
				_, exist := nums[k]
				if !exist {
					nums[k] = true
					sum += k
				}
			}
		}
	}
	return sum
}
