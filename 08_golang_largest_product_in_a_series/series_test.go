package series

import (
	"testing"

	"gotest.tools/assert"
)

func TestLargetProduct(t *testing.T) {
	assert.Equal(t, 9*9*8, largetProduct(3))
	assert.Equal(t, 5832, largetProduct(4))
	assert.Equal(t, 23514624000, largetProduct(13))
}
