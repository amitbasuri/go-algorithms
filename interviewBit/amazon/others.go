package amazon

func majorityElement(A []int) int {
	s := len(A) / 2
	m := make(map[int]int)
	for _, a := range A {
		m[a]++
		if m[a] > s {
			return a
		}
	}
	return 0
}

func minIntSliceIndex(arr []int) int {
	var minIndex int
	for i, a := range arr {
		if a < arr[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

//Distribute Candy
//There are N children standing in a line. Each child is assigned a rating value.
//
//You are giving candies to these children subjected to the following requirements:
//
//1. Each child must have at least one candy.
//2. Children with a higher rating get more candies than their neighbors.
func candy(A []int) int {
	if len(A) == 0 {
		return 0
	}
	minIndex := minIntSliceIndex(A)
	candy := make([]int, len(A))
	candy[minIndex] = 1
	numOfCandies := 1
	for index := minIndex + 1; index < len(A); index++ {
		if A[index] > A[index-1] {
			candy[index] = candy[index-1] + 1
		} else {
			candy[index] = 1
		}
		numOfCandies += candy[index]
	}

	for index := minIndex - 1; index >= 0; index-- {
		if A[index] > A[index+1] {
			candy[index] = candy[index+1] + 1
		} else {
			candy[index] = 1
		}
		numOfCandies += candy[index]
	}
	return numOfCandies
}

func maxProductSubarray(A []int) int {

	maxP := A[0]
	minP := A[0]
	res := A[0]
	for i := 1; i < len(A); i++ {
		cMaxP := maxInt(maxP*A[i], minP*A[i], A[i])
		cMinP := minInt(maxP*A[i], minP*A[i], A[i])
		if cMaxP > res {
			res = cMaxP
		}
		maxP = cMaxP
		minP = cMinP
	}

	return res
}

func maxMinProductSubarraySequence(arr []int, n int) (int, int) {

	if n == 1 {
		return arr[n-1], arr[n-1]
	}

	max, min := maxMinProductSubarraySequence(arr, n-1)
	if max == 0 {
		max = 1
	}
	if min == 0 {
		min = 1
	}
	return maxInt(arr[n-1]*max, arr[n-1]*min, max), minInt(arr[n-1]*max, arr[n-1]*min, min)
}
