package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCandy(t *testing.T) {
	assert.Equal(t, 1, candy([]int{2}))
	assert.Equal(t, 3, candy([]int{2, 3}))
	assert.Equal(t, 3, candy([]int{3, 2}))
	assert.Equal(t, 4, candy([]int{2, 3, 2}))

}

func TestMaxProductSubarray(t *testing.T) {
	assert.Equal(t, 10000, maxProductSubarray([]int{10, -10, 10, -10}))
	assert.Equal(t, -4, maxProductSubarray([]int{-4}))

	A := []int{0, 3}
	assert.Equal(t, 3, maxProductSubarray(A))
	A = []int{3, 0}
	assert.Equal(t, 3, maxProductSubarray(A))
	A = []int{-3, 0}
	assert.Equal(t, 0, maxProductSubarray(A))

	A = []int{0, -3, 4, -10, -1, -6, 0, 8, -8, -6, -5, -5, 0, -3, -9, 1, 5, -8, 0, 6, 1, -6, -8, 3, 0, -8, -9, 6, 8, 5}
	assert.Equal(t, 17280, maxProductSubarray(A))
}
