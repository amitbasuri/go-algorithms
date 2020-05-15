package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSubsetSumRecursive(t *testing.T) {
	b := countSubsetSumRecursive([]int{1, 2, 3, 3}, 4, 6)
	assert.Equal(t, 3, b)

	c := countSubsetSumRecursive([]int{1, 1, 1, 1}, 4, 1)
	assert.Equal(t, 4, c)
}

func TestCountSubsetSum(t *testing.T) {
	b := countSubsetSum([]int{1, 2, 3, 3}, 4, 6)
	assert.Equal(t, 3, b)

	c := countSubsetSum([]int{1, 1, 1, 1}, 4, 1)
	assert.Equal(t, 4, c)
}
