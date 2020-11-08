package interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestSubarrayWithSum(t *testing.T) {
	assert.Equal(t, 2, ShortestSubarrayWithSum([]int{1, 2, 3, 4}, 6))
	assert.Equal(t, 4, ShortestSubarrayWithSum([]int{1, 2, 3, 4}, 10))
	assert.Equal(t, 1, ShortestSubarrayWithSum([]int{11, 21, 3, 4}, 20))
	assert.Equal(t, 1, ShortestSubarrayWithSum([]int{17, 25, 34, 4}, 3))
	assert.Equal(t, -1, ShortestSubarrayWithSum([]int{11, 2, 3, 4}, 303))
}
