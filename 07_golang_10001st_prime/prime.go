package prime

func nthPrime(n int) int {
	counter := 0
	i := 2
	for counter != n {
		if isPrime(i) {
			counter++
		}
		i++
	}
	return i - 1
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
