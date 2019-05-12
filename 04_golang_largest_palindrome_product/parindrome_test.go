package parindrome

import (
	"testing"

	"gotest.tools/assert"
)

func TestIsParindrome(t *testing.T) {
	assert.Assert(t, isParindrome(123454321))
	assert.Assert(t, isParindrome(334433))
	assert.Assert(t, !isParindrome(3344337))
	assert.Assert(t, !isParindrome(12))
}

func TestLargesetParindromeProduct(t *testing.T) {
	assert.Equal(t, 9009, largesetParindromeProduct(2))
	assert.Equal(t, 906609, largesetParindromeProduct(3))
}
