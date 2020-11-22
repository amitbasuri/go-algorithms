package codility

// max profit stocks
func MaxProfitSolution(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	var maxProfit int

	for _, price := range prices {

		if minPrice < price { // sell

			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}

		if price < minPrice {
			minPrice = price
		}

	}

	return maxProfit

}

func MaxDoubleSliceSumSolution(arr []int) int {

	maxSum := 0
	minElem := arr[1]

	runningSum := 0
	for i := 1; i < len(arr)-1; i++ {
		runningSum += arr[i]
		if runningSum > maxSum {
			maxSum = runningSum
		}

		if arr[i] < minElem {
			minElem = arr[i]
		}
	}

	return maxSum - minElem

}
