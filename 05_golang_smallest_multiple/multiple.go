package multiple

import "math"

func smallestMultiple(max int) int {
	factors := make(map[int]int)
	for i := 2; i <= max; i++ {
		tempFactors := factorize(i)
		for prime, power := range tempFactors {
			if v, ok := factors[prime]; ok {
				if v < power {
					factors[prime] = power
				}
			} else {
				factors[prime] = power
			}
		}
	}
	return defactorize(factors)
}

func defactorize(factors map[int]int) int {
	ans := 1
	for prime, power := range factors {
		ans *= int(math.Pow(float64(prime), float64(power)))
	}
	return ans
}

func factorize(in int) map[int]int {
	factors := make(map[int]int)
	counter := 2
	for !isPrime(in) {
		if isPrime(counter) {
			if in%counter == 0 {
				in = in / counter
				factors[counter]++
			} else {
				counter++
			}
		} else {
			counter++
		}
	}
	factors[in]++
	return factors
}

// AKS primality test
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
