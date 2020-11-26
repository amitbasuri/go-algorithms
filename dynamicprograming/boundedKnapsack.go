package dynamicprograming

//0-1 Knapsack Problem | DP-10
//Given weights and values of n items, put these items in a knapsack of capacity W to get the maximum
//total value in the knapsack. In other words, given two integer arrays val[0..n-1] and wt[0..n-1]
//which represent values and weights associated with n items respectively. Also given an integer W
//which represents knapsack capacity, find out the maximum value subset of val[] such that sum of
//the weights of this subset is smaller than or equal to W. You cannot break an item, either pick
//the complete item or don’t pick it (0-1 property).
func KnapsackTD(value []int, weight []int, weightLeft int, size int) int {
	if weightLeft == 0 || size == 0 {
		return 0
	}

	if weight[size-1] > weightLeft {
		return KnapsackTD(value, weight, weightLeft, size-1)
	}

	taken := KnapsackTD(value, weight, weightLeft-weight[size-1], size-1) + value[size-1]
	notTaken := KnapsackTD(value, weight, weightLeft, size-1)

	if notTaken > taken {
		return notTaken
	}
	return taken
}

func KnapsackBU(value []int, weight []int, weightLeft int, size int) int {
	dp := make([][]int, size+1)
	for sizeIndex := 0; sizeIndex < size+1; sizeIndex++ {
		dp[sizeIndex] = make([]int, weightLeft+1)
	}

	for sizeIndex := 1; sizeIndex <= size; sizeIndex++ {
		for weightLeftIndex := 1; weightLeftIndex <= weightLeft; weightLeftIndex++ {
			if weight[sizeIndex-1] > weightLeftIndex {
				dp[sizeIndex][weightLeftIndex] = dp[sizeIndex-1][weightLeftIndex]
			} else {
				taken := dp[sizeIndex-1][weightLeftIndex-weight[sizeIndex-1]] + value[sizeIndex-1]
				notTaken := dp[sizeIndex-1][weightLeftIndex]
				if taken > notTaken {
					dp[sizeIndex][weightLeftIndex] = taken
				} else {
					dp[sizeIndex][weightLeftIndex] = notTaken
				}
			}
		}
	}

	return dp[size][weightLeft]

}
