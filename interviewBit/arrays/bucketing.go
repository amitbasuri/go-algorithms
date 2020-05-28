package arrays

import (
	"sort"
)

//Max Distance
//Asked in: Google Amazon Microsoft
//Given an array A of integers, find the maximum of j - i subjected to the constraint of A[i] <= A[j].
//
//If there is no solution possible, return -1.
//
//Example :
//A : [3 5 4 2]
//Output : 2
//for the pair (3, 4)
func maximumGap(arr []int) int {

	//  not passing all test cases
	// correct solution : max(RightMinArr[i], leftMinArr[i])

	indexMap := make(map[int]int)
	for i, a := range arr {
		if _, ok := indexMap[a]; !ok {
			indexMap[a] = i
		}
	}
	sort.Ints(arr)

	max := -1 << 63
	for i, a := range arr {
		diff := (len(arr) - 1 - i) - indexMap[a]
		if max < diff {
			max = diff
		}
	}
	return max
}

//Wave Array
//Asked in: Google Adobe Amazon
//Given an array of integers, sort the array into a wave like array and return it,
//In other words, arrange the elements into a sequence such that a1 >= a2 <= a3 >= a4 <= a5.....
//
//Example
//Given [1, 2, 3, 4]
//One possible answer : [2, 1, 4, 3]
//Another possible answer : [4, 1, 3, 2]
//NOTE : If there are multiple answers possible, return the one that is lexicographically smallest.
//So, in example case, you will return [2, 1, 4, 3]
func wave(arr []int) []int {
	sort.Ints(arr)
	i := 0
	for i+1 < len(arr) {
		arr[i], arr[i+1] = arr[i+1], arr[i]
		i += 2
	}
	return arr
}

//Maximum Unsorted Subarray
//Asked in: Amazon
//You are given an array (zero indexed) of N non-negative integers, A0, A1 ,…, AN-1.
//Find the minimum sub array Al, Al+1 ,…, Ar so if we sort(in ascending order) that sub array,
//then the whole array should get sorted.
//If A is already sorted, output -1.
//Example :
//Input 1:
//A = [1, 3, 2, 4, 5]
//Return: [1, 2]
//Input 2: A = [1, 2, 3, 4, 5]
//Return: [-1]
func subUnsort(arr []int) []int {
	if sort.IntsAreSorted(arr) {
		return []int{-1}
	}
	left := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			left = i - 1
			break
		}
	}

	right := len(arr) - 1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			right = i + 1
			break
		}
	}

	middle := make([]int, right-left+1)
	copy(middle, arr[left:right+1])
	sort.Ints(middle)
	leftVal := middle[0]
	rightVal := middle[len(middle)-1]

	for i := left - 1; i > 0; i-- {
		if arr[i] <= leftVal {
			break
		}
		left = i
	}

	for i := right + 1; i < len(arr); i++ {
		if arr[i] >= rightVal {
			break
		}
		right = i
	}
	return []int{left, right}

}
