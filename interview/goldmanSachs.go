package interview

//Shortest Subarray with Sum at Least K
//
//Return the length of the shortest, non-empty,  subarray of A with sum at least K.
//If there is no non-empty subarray with sum at least K, return -1.
//

func ShortestSubarrayWithSum(arr []int, x int) int {
	s := ShortestSubarrayWithSumHelper(arr, len(arr)-1, x)
	if s > len(arr) {
		return -1
	}
	return s
}

func ShortestSubarrayWithSumHelper(arr []int, i int, x int) int {

	if x <= 0 {
		return 0
	}

	if i < 0 {
		return len(arr) + 1
	}

	taken := ShortestSubarrayWithSumHelper(arr, i-1, x)
	notTaken := ShortestSubarrayWithSumHelper(arr, i-1, x-arr[i]) + 1

	return min(taken, notTaken)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
