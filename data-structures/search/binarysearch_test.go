package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {

	t.Run("has value", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8}

		assert.True(t, BinarySearch(array, 1))
		assert.True(t, BinarySearch(array, 2))
		assert.True(t, BinarySearch(array, 3))
		assert.True(t, BinarySearch(array, 4))
		assert.True(t, BinarySearch(array, 5))
		assert.True(t, BinarySearch(array, 6))
		assert.True(t, BinarySearch(array, 7))
		assert.True(t, BinarySearch(array, 8))

		assert.False(t, BinarySearch(array, 0))
		assert.False(t, BinarySearch(array, 9))
	})
}
