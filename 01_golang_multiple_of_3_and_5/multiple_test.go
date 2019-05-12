package multiple

import (
	"testing"

	"gotest.tools/assert"
)

func TestMultiple(t *testing.T) {
	assert.Equal(t, 23, Multiple(10, 3, 5))
	assert.Equal(t, 0, Multiple(3, 3, 5))
	assert.Equal(t, 45, Multiple(15, 3, 5))
	assert.Equal(t, 34, Multiple(15, 7, 13))
	//answer
	assert.Equal(t, 233168, Multiple(1000, 3, 5))
}
