package parindrome

import (
	"math"
	"strconv"
)

func isParindrome(in int) bool {
	s := strconv.Itoa(in)
	reverse := ""
	for _, char := range s {
		reverse = string(char) + reverse
	}
	return reverse == s
}

func largesetParindromeProduct(digit int) int {
	max := 0
	counter := int(math.Pow(float64(10), float64(digit)))
	for i := 0; i < counter; i++ {
		for j := 0; j < counter; j++ {
			tmp := i * j
			if isParindrome(tmp) && tmp > max {
				max = tmp
			}
		}
	}
	return max
}
