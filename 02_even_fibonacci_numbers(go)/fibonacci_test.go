package fibonacci

import (
	"testing"

	"gotest.tools/assert"
)

func TestFibonacci(t *testing.T) {
	assert.DeepEqual(t, []int{1}, Fibonacci(1))
	assert.DeepEqual(t, []int{1, 2}, Fibonacci(2))
	assert.DeepEqual(t, []int{1, 2, 3}, Fibonacci(3))
	assert.DeepEqual(t, []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}, Fibonacci(10))
}

func TestAnwser(t *testing.T) {
	assert.Equal(t, 0, evenSumFib(1))
	assert.Equal(t, 2, evenSumFib(2))
	assert.Equal(t, 2, evenSumFib(3))
	assert.Equal(t, 10, evenSumFib(8))
	assert.Equal(t, 44, evenSumFib(35))
	// answer
	assert.Equal(t, 4613732, evenSumFib(4000000))
}
