package prime

import (
	"testing"

	"gotest.tools/assert"
)

func TestNthPrime(t *testing.T) {
	assert.Equal(t, 2, nthPrime(1))
	assert.Equal(t, 5, nthPrime(3))
	assert.Equal(t, 104743, nthPrime(10001)) // 10000th prime
}
