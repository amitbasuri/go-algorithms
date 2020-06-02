package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNexGreaterElementNaive(t *testing.T) {
	assert.Equal(t, []int{2, -1}, NextGreaterElementNaive([]int{1, 2}))
	assert.Equal(t, []int{5, 25, 25, -1}, NextGreaterElementNaive([]int{4, 5, 2, 25}))
}

func TestNexGreaterElement(t *testing.T) {
	assert.Equal(t, []int{2, -1}, NextGreaterElement([]int{1, 2}))
	assert.Equal(t, []int{-1, -1}, NextGreaterElement([]int{2, 1}))
	assert.Equal(t, []int{5, 25, 25, -1}, NextGreaterElement([]int{4, 5, 2, 25}))
}

func TestPreviousGreaterElement(t *testing.T) {
	assert.Equal(t, []int{-1, -1}, PreviousGreaterElement([]int{1, 2}))
	assert.Equal(t, []int{-1, -1, 5, -1}, PreviousGreaterElement([]int{4, 5, 2, 25}))
}

func TestPreviousSmallerElement(t *testing.T) {
	assert.Equal(t, []int{-1, 1}, PreviousSmallerElement([]int{1, 2}))
	assert.Equal(t, []int{-1, -1}, PreviousSmallerElement([]int{2, 1}))
	assert.Equal(t, []int{-1, 4, -1, 2}, PreviousSmallerElement([]int{4, 5, 2, 25}))
}

func TestStockSpanProblem(t *testing.T) {
	assert.Equal(t, []int{1, 1, 1, 2, 1, 4, 6}, StockSpanProblem([]int{100, 80, 60, 70, 60, 75, 85}))
	assert.Equal(t, []int{1, 2, 1, 4}, StockSpanProblem([]int{4, 5, 2, 25}))
}

func TestMaximumAreaHistogram(t *testing.T) {
	assert.Equal(t, 12, MaximumAreaHistogram([]int{6, 2, 5, 4, 5, 1, 6}))
	assert.Equal(t, 3, MaximumAreaHistogram([]int{1, 2, 1}))
	assert.Equal(t, 4, MaximumAreaHistogram([]int{1, 2, 2}))

}
