package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetSum(t *testing.T) {
	b := subSetSumWithoutMemo([]int{1, 2, 4, 7, 9}, 5, 19)
	assert.True(t, b)

	c := subSetSumWithoutMemo([]int{13, 2, 4, 7, 9}, 5, 5)
	assert.False(t, c)
}

func BenchmarkSubSetSumWithoutMemo(b *testing.B) {
	arr := []int{}
	for i := 1; i <= 25; i++ {
		arr = append(arr, i)
	}
	for n := 0; n < b.N; n++ {
		subSetSumWithoutMemo(arr, len(arr), 19)
	}

}

func TestSubsetSumWithMemo(t *testing.T) {
	m := make(subsetSumMemo)
	b := m.subsetSum([]int{1, 2, 4, 7, 9}, 5, 19)
	assert.True(t, b)

	m = make(subsetSumMemo)

	c := m.subsetSum([]int{13, 2, 4, 7, 9}, 5, 5)
	assert.False(t, c)
}

func BenchmarkSubSetSumWithMemo(b *testing.B) {
	arr := []int{}
	for i := 1; i <= 25; i++ {
		arr = append(arr, i)
	}
	for n := 0; n < b.N; n++ {
		m := make(subsetSumMemo)
		m.subsetSum(arr, len(arr), 390)
	}

}

func TestSubsetSumWithTD(t *testing.T) {
	b := subsetSumWithTD([]int{1, 2, 4, 7, 9}, 5, 19)
	assert.True(t, b)

	c := subsetSumWithTD([]int{13, 2, 4, 7, 9}, 5, 5)
	assert.False(t, c)
}

func BenchmarkSubSetSumWithTD(b *testing.B) {
	arr := []int{}
	for i := 1; i <= 25; i++ {
		arr = append(arr, i)
	}
	for n := 0; n < b.N; n++ {
		subsetSumWithTD(arr, len(arr), 390)
	}

}
