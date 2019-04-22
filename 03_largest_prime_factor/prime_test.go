package prime

import (
	"testing"

	"gotest.tools/assert"
)

func TestPrimeFactor(t *testing.T) {
	assert.DeepEqual(t, []int{5, 7, 13, 29}, primeFactors(13195))
	// answer
	assert.DeepEqual(t, []int{5, 7, 13, 29, 71, 839, 1471, 6857}, primeFactors(600851475143))
}

func TestIsPrime(t *testing.T) {
	assert.Assert(t, isPrime(29))
	assert.Assert(t, isPrime(104729)) // 10000th prime
}
