package amazon

//Best Time to Buy and Sell Stocks I
//If you were only permitted to complete at most one transaction
//(ie, buy one and sell one share of the stock), design an algorithm to find the maximum profit.
//A = [1, 4, 5, 2, 4] => 4; 5-1
//A = [3, 4, 5, 2, 4] => 2

// find max(arr[i]-arr[j]) where i > j
func maxProfit(A []int) int {
	min := 0
	maxSoFar := make([]int, len(A))
	for i, v := range A {
		if A[min] > v {
			min = i //change min val index
		}
		maxSoFar[i] = v - A[min]
		if i > 0 && maxSoFar[i] < maxSoFar[i-1] {
			maxSoFar[i] = maxSoFar[i-1]
		}
	}
	if len(maxSoFar) == 0 {
		return 0
	}
	return maxSoFar[len(maxSoFar)-1]
}

//Best Time to Buy and Sell Stocks II
//
//You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
//
//However, you may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

//Input 1:
//A = [1, 2, 3]
//
//Output 1:
//2
//
//Explanation 1:
//=> Buy a stock on day 0.
//=> Sell the stock on day 1. (Profit +1)
//=> Buy a stock on day 1.
//=> Sell the stock on day 2. (Profit +1)
//
//Overall profit = 2

func maxProfit2(A []int) int {
	sum := 0
	for i := 0; i < len(A)-1; i++ {
		if A[i+1]-A[i] > 0 { // sell if price inc from last time
			sum += A[i+1] - A[i]
		}
	}
	return sum
}

//Design an algorithm to find the maximum profit. You may complete at most 2 transactions.
//Return the maximum possible profit.
//Note: You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
//Input 1:
//A = [1, 2, 1, 2,3]
//Output 1: 2
//Explanation 1:
//Day 0 : Buy
//Day 1 : Sell
//Day 2 : Buy
//Day 3 : Sell
//
//Input 2:
//A = [7, 2, 4, 8, 7]
//Output 2: 6
//Explanation 2:
//Day 1 : Buy
//Day 3 : Sell
func maxProfit3(arr []int) int {
	sum := 0
	if len(arr) < 2 {
		return 0
	}
	buy := arr[0]
	n := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			sum += arr[i] - arr[i-1]
			continue

		}
		buy = arr[i]
		n++
		if n == 2 {
			break
		}

	}
	buy = 0
	return sum + buy
}
