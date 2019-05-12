package multiple

import (
	"testing"

	"gotest.tools/assert"
)

func TestSmallestMultiple(t *testing.T) {
	assert.Equal(t, 2520, smallestMultiple(10))
	assert.Equal(t, 232792560, smallestMultiple(20))
}

func TestFactorize(t *testing.T) {
	assert.DeepEqual(t, map[int]int{2: 1, 5: 1}, factorize(10))
	assert.DeepEqual(t, map[int]int{2: 2, 5: 1}, factorize(20))
	assert.DeepEqual(t, map[int]int{2: 2, 3: 2}, factorize(36))

	assert.DeepEqual(t, map[int]int{3: 2, 17: 2, 379721: 1}, factorize(987654321))
	assert.DeepEqual(t, map[int]int{17: 1}, factorize(17))
	assert.DeepEqual(t, map[int]int{2: 4}, factorize(16))

	assert.DeepEqual(t, map[int]int{2: 2, 5: 2}, factorize(100))
}
func TestDefactorize(t *testing.T) {
	assert.Equal(t, defactorize(map[int]int{2: 1, 5: 1}), 10)
	assert.Equal(t, defactorize(map[int]int{2: 2, 3: 2}), 36)

	assert.Equal(t, defactorize(map[int]int{3: 2, 17: 2, 379721: 1}), 987654321)
}
