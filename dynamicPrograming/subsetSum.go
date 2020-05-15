package dynamicPrograming

//Given a set of non-negative integers, and a value sum, determine if there is a subset of the given set
//with sum equal to given sum.
//
//Example:
//
//Input: set[] = {3, 34, 4, 12, 5, 2}, sum = 9
//Output: True
//Explanation: There is a subset (4, 5) with sum 9.

type sunsetSumTopDown [][]bool

func subsetSumWithTD(arr []int, n int, sum int) bool {
	td := make(sunsetSumTopDown, n+1)

	for i := 0; i <= n; i++ {
		td[i] = make([]bool, sum+1)
		td[i][0] = true
	}

	return td.subsetSum(arr, n, sum)
}

func (td sunsetSumTopDown) subsetSum(arr []int, n int, sum int) bool {

	for currLen := 1; currLen <= n; currLen++ {
		for currSum := 1; currSum <= sum; currSum++ {
			currN := arr[currLen-1]
			if currN <= currSum {
				taken := td[currLen-1][currSum-currN]
				notTaken := td[currLen-1][currSum]

				td[currLen][currSum] = taken || notTaken

				continue
			}

			notTaken := td[currLen-1][currSum]

			td[currLen][currSum] = notTaken
		}
	}
	return td[n][sum]
}

type subsetSumMemo map[[2]int]bool

func (m subsetSumMemo) subsetSum(arr []int, n int, sum int) bool {
	if sum == 0 {
		return true
	}

	if n == 0 {
		return false
	}
	var taken, notTaken bool

	if arr[n-1] <= sum {
		if v, ok := m[[2]int{n - 1, sum - arr[n-1]}]; ok {
			taken = v
		} else {
			taken = m.subsetSum(arr, n-1, sum-arr[n-1])
			m[[2]int{n - 1, sum - arr[n-1]}] = taken
		}

		if v, ok := m[[2]int{n - 1, sum}]; ok {
			notTaken = v
		} else {
			notTaken = m.subsetSum(arr, n-1, sum)
			m[[2]int{n - 1, sum}] = notTaken

		}

		m[[2]int{n, sum}] = taken || notTaken
		return taken || notTaken

	}

	if v, ok := m[[2]int{n - 1, sum}]; ok {
		notTaken = v
	} else {
		notTaken = m.subsetSum(arr, n-1, sum)
		m[[2]int{n - 1, sum}] = notTaken
	}

	m[[2]int{n, sum}] = notTaken
	return notTaken
}

func subSetSumWithoutMemo(arr []int, n int, sum int) bool {
	if sum == 0 {
		return true
	}
	if n == 0 {
		return false
	}

	taken := subSetSumWithoutMemo(arr, n-1, sum-arr[n-1])
	notTaken := subSetSumWithoutMemo(arr, n-1, sum)

	return taken || notTaken

}
