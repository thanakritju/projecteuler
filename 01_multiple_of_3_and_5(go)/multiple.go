package multiple

import "fmt"

// Multiple returns sum of all the natural numbers below 'n' that are multiples of 'a' or 'b'
func Multiple(n, a, b int) int {
	var sum int
	if a == 1 || b == 1 {
		fmt.Println("This function does not calculate multiples of 1")
		return 0
	}
	for i := 1; i < n; i++ {
		if i%a == 0 || i%b == 0 {
			sum += i
		}
	}
	return sum
}
