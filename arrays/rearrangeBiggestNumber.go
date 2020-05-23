package arrays

import (
	"sort"
	"strconv"
)

//Given an array of numbers, arrange them in a way that yields the largest value. For example,
//if the given numbers are {54, 546, 548, 60}, the arrangement 6054854654 gives the largest value.
//	And if the given numbers are {1, 34, 3, 98, 9, 76, 45, 4},
//then the arrangement 998764543431 gives the largest value.

type RearrangedLargestNumberArr []int

// Len is the number of elements in the collection.
func (arr RearrangedLargestNumberArr) Len() int {
	return len(arr)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (arr RearrangedLargestNumberArr) Less(i, j int) bool {
	// > for desc order, < for asc order
	return strconv.Itoa(arr[i])+strconv.Itoa(arr[j]) > strconv.Itoa(arr[j])+strconv.Itoa(arr[i])
}

// Swap swaps the elements with indexes i and j.
func (arr RearrangedLargestNumberArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func RearrangedLargestNumber(arr []int) int {
	a := RearrangedLargestNumberArr(arr)
	sort.Sort(a)
	s := ""
	for _, val := range a {

		s += strconv.Itoa(val)
	}
	sol, _ := strconv.Atoi(s)
	return sol
}
