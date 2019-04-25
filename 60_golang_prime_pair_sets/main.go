package main

import "strconv"

func main() {
	// TODO
}

func icp(num1 int, num2 int) bool {
	cat1, _ := strconv.ParseInt(string(num1)+string(num2), 10, 64)
	cat2, _ := strconv.ParseInt(string(num2)+string(num1), 10, 64)
	return isPrime(cat1) && isPrime(cat2)
}

func generatePrimes(n int) []int64 {
	count := 0
	primes := make([]int64, n)
	for i := int64(2); count < n; i++ {
		if isPrime(i) {
			primes[count] = i
			count++
		}
	}
	return primes
}

func isPrime(n int64) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	var i int64 = 5
	for i*i < n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}
