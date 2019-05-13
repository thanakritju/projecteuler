package sumsquare

import (
	"testing"

	"gotest.tools/assert"
)

func TestSumSquare(t *testing.T) {
	assert.Equal(t, 2640, sumSquare(10))
	assert.Equal(t, 25164150, sumSquare(100))
}
