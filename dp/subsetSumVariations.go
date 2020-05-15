package dp

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
