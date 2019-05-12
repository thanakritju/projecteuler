package prime

func primeFactors(value int) []int {
	factors := make([]int, 0)
	for i := 2; i <= value; i++ {
		if value%i == 0 && isPrime(i) {
			value = value / i
			factors = append(factors, i)
		}
	}
	return factors
}

// AKS primality test
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}
