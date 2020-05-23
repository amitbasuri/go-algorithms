package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRearrangedLargestNumber(t *testing.T) {

	assert.Equal(t, 9876, RearrangedLargestNumber([]int{7, 9, 6, 8}))
	assert.Equal(t, 98773, RearrangedLargestNumber([]int{9, 8, 77, 3}))

}

func TestMergeSortedArray(t *testing.T) {
	assert.Equal(t, []int{}, mergeSortedArrays([]int{}, []int{}))
	assert.Equal(t, []int{0, 1}, mergeSortedArrays([]int{0}, []int{1}))
	assert.Equal(t, []int{3, 5, 7, 8, 9}, mergeSortedArrays([]int{3, 5, 7}, []int{8, 9}))
	assert.Equal(t, []int{3, 4, 5, 5, 7, 8, 8}, mergeSortedArrays([]int{3, 5, 7, 8}, []int{4, 5, 8}))
	assert.Equal(t, []int{3, 5, 7, 8, 9}, mergeSortedArrays([]int{8, 9}, []int{3, 5, 7}))

}
