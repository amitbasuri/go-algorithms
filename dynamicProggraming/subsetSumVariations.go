package dynamicProggraming

//Partition problem is to determine whether a given set can be partitioned into two subsets
//such that the sum of elements in both subsets is same.
//Examples:
//
//arr[] = {1, 5, 11, 5}
//Output: true
//The array can be partitioned as {1, 5, 5} and {11}
//
//arr[] = {1, 5, 3}
//Output: false
//The array cannot be partitioned into equal sum sets.

func equalSumPartition(arr []int, n int) bool {
	var sum int
	for _, a := range arr {
		sum += a
	}
	if sum%2 == 1 {
		return false
	}

	return subsetSumWithTD(arr, n, sum/2)
}

//Given an array arr[] of length N and an integer X, the task is to find the number of subsets with sum equal to X.
//
//Examples:
//
//Input: arr[] = {1, 2, 3, 3}, X = 6
//Output: 3
//All the possible subsets are {1, 2, 3}, {1, 2, 3} and {3, 3}
//
//Input: arr[] = {1, 1, 1, 1}, X = 1
//Output: 4

func countSubsetSumRecursive(arr []int, n int, subsetSum int) int {
	if subsetSum == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}

	choice := arr[n-1]
	if choice <= subsetSum {
		taken := countSubsetSumRecursive(arr, n-1, subsetSum-choice)
		notTaken := countSubsetSumRecursive(arr, n-1, subsetSum)
		return taken + notTaken
	}

	notTaken := countSubsetSumRecursive(arr, n-1, subsetSum)
	return notTaken
}

type countSubsetSumTopDown [][]int

func countSubsetSum(arr []int, n int, sum int) int {

	td := make(countSubsetSumTopDown, n+1)

	for i := 0; i <= n; i++ {
		td[i] = make([]int, sum+1)
		td[i][0] = 1
	}
	td.Fill(arr, n, sum)
	return td[n][sum]
}

func (td countSubsetSumTopDown) Fill(arr []int, n int, sum int) {

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			choice := arr[i-1]

			if choice <= j {
				td[i][j] = td[i-1][j-choice] + td[i-1][j]
				continue
			}
			td[i][j] = td[i-1][j]
		}
	}
}
