package binarySearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchAsc(t *testing.T) {
	assert.Equal(t, 2, SearchAsc([]int{1, 3, 4, 6, 7, 9, 12, 15, 16}, 4))
	assert.Equal(t, -1, SearchAsc([]int{1, 3, 4, 6, 7, 9, 12, 15, 16}, 19))
	assert.Equal(t, -1, SearchAsc([]int{1, 3, 4, 6, 7, 9, 12, 15, 16}, 8))
	assert.Equal(t, -1, SearchAsc([]int{1}, 8))
	assert.Equal(t, 0, SearchAsc([]int{1}, 1))
}

func TestSearchDesc(t *testing.T) {
	assert.Equal(t, 2, SearchDesc([]int{12, 10, 9, 8, 6, 4, 3, 2}, 9))
	assert.Equal(t, -1, SearchDesc([]int{12, 10, 9, 8, 6, 4, 3, 2}, 19))
	assert.Equal(t, -1, SearchDesc([]int{12, 10, 9, 8, 6, 4, 3, 2}, 7))
	assert.Equal(t, -1, SearchDesc([]int{1}, 8))
	assert.Equal(t, 0, SearchDesc([]int{1}, 1))
}

func TestFirstOccurance(t *testing.T) {
	assert.Equal(t, 2, FirstOccurance([]int{1, 3, 6, 6, 6, 6, 6, 15, 16}, 6))
	assert.Equal(t, -1, FirstOccurance([]int{1, 3, 4, 6, 7, 9, 12, 15, 16}, 19))
	assert.Equal(t, 6, FirstOccurance([]int{1, 3, 4, 6, 7, 9, 12, 12, 16}, 12))
	assert.Equal(t, -1, FirstOccurance([]int{1}, 8))
	assert.Equal(t, 0, FirstOccurance([]int{1}, 1))
}

func TestRotationCount(t *testing.T) {
	assert.Equal(t, 0, RotationCount([]int{1, 4, 5, 6, 7, 8, 9, 13}))
	assert.Equal(t, 1, RotationCount([]int{4, 5, 6, 7, 8, 9, 13, 2}))
	assert.Equal(t, 7, RotationCount([]int{13, 1, 4, 5, 6, 7, 8, 9}))
}

func TestRotatedArrSearch(t *testing.T) {
	arr := []int{192, 194, 197, 198, 199, 200, 201, 203, 204, 2, 3, 4, 7, 9, 10,
		11, 14, 15, 16, 17, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 33, 34, 35,
		36, 39, 40, 42, 43, 46, 47, 50, 51, 52, 53, 55, 57, 59, 60, 64, 66, 68, 70,
		71, 72, 75, 76, 78, 85, 86, 87, 91, 92, 94, 95, 96, 99, 102, 105, 106, 107,
		109, 111, 113, 114, 115, 119, 120, 121, 123, 125, 129, 130, 131, 133, 134,
		137, 138, 139, 140, 141, 142, 143, 145, 146, 149, 151, 152, 156, 160, 161,
		165, 167, 168, 170, 171, 176, 177, 178, 179, 180, 181, 182, 185, 186, 190}
	assert.Equal(t, 52, RotatedArrSearch(arr, 70))
}

func TestNearlySortedArrSearch(t *testing.T) {
	assert.Equal(t, 2, NearlySortedArrSearch([]int{10, 3, 40, 20, 50, 80, 70}, 40))
	assert.Equal(t, 7, NearlySortedArrSearch([]int{2, 3, 10, 4, 40, 50, 62, 60, 65}, 60))
	assert.Equal(t, 0, NearlySortedArrSearch([]int{10, 3, 40, 20, 50, 80, 70}, 10))
	assert.Equal(t, -1, NearlySortedArrSearch([]int{10, 3, 40, 20, 50, 80, 70}, 90))
	assert.Equal(t, -1, NearlySortedArrSearch([]int{10, 3, 40, 20, 50, 80, 70}, 2))
}

func TestFloorInSortedArr(t *testing.T) {
	assert.Equal(t, 40, FloorInSortedArr([]int{10, 20, 30, 40}, 40))
	assert.Equal(t, 10, FloorInSortedArr([]int{10, 20, 30, 40}, 10))
	assert.Equal(t, 20, FloorInSortedArr([]int{10, 20, 30, 40}, 25))
	assert.Equal(t, 30, FloorInSortedArr([]int{10, 20, 30, 40, 60}, 35))
	assert.Equal(t, -1, FloorInSortedArr([]int{10, 20, 30, 40, 60}, 5))
}

func TestNextLetter(t *testing.T) {
	assert.Equal(t, "b", NextLetter("abcde", "b"))
	assert.Equal(t, "c", NextLetter("acdexz", "b"))
	assert.Equal(t, "c", NextLetter("acde", "b"))
	assert.Equal(t, "", NextLetter("abcde", "f"))

}
