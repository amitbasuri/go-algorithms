package dynamicPrograming

//Rod Cutting Problem
//Given a rod of length n inches and an array of prices that contains prices of
//all pieces of size smaller than n. Determine the  maximum value obtainable
//by cutting up the rod and selling the pieces.
//Example:
//if length of the rod is 8 and the values of different pieces are given
//as following, then the maximum obtainable value is 22 (by cutting in
//two pieces of lengths 2 and 6)
//
//length   | 1   2   3   4   5   6   7   8
//--------------------------------------------
//price    | 1   5   8   9  10  17  17  20

func RodCuttingProblem(length []int, price []int, rodSize int, optionsSize int) int {
	if rodSize <= 0 || optionsSize == 0 {
		return 0
	}

	if rodSize < length[optionsSize-1] {
		return RodCuttingProblem(length, price, rodSize, optionsSize-1)
	}

	cut := price[optionsSize-1] + RodCuttingProblem(length, price, rodSize-length[optionsSize-1], optionsSize)
	notCut := RodCuttingProblem(length, price, rodSize, optionsSize-1)

	if cut > notCut {
		return cut
	}
	return notCut

}

func RodCuttingProblemBU(length []int, price []int, rodSize int, optionsSize int) int {
	dp := make([][]int, rodSize+1)
	for i := 0; i <= rodSize; i++ {
		dp[i] = make([]int, optionsSize+1)
	}

	for rodSizeIndex := 1; rodSizeIndex <= rodSize; rodSizeIndex++ {
		for optionsSizeIndex := 1; optionsSizeIndex <= optionsSize; optionsSizeIndex++ {

			if rodSizeIndex < length[optionsSizeIndex-1] {
				dp[rodSizeIndex][optionsSizeIndex] = dp[rodSizeIndex][optionsSizeIndex-1]
				continue
			}

			cut := price[optionsSizeIndex-1] + dp[rodSizeIndex-length[optionsSizeIndex-1]][optionsSizeIndex]
			notCut := dp[rodSizeIndex][optionsSizeIndex-1]
			if cut > notCut {
				dp[rodSizeIndex][optionsSizeIndex] = cut
			} else {
				dp[rodSizeIndex][optionsSizeIndex] = notCut
			}
		}
	}

	return dp[rodSize][optionsSize]

}

//Coin Change Problem Maximum Number of ways
//Given a value N, if we want to make change for N cents, and we have infinite
//supply of each of S = { S1, S2, .. , Sm} valued coins, how many ways can we make
//the change? The order of coins doesn’t matter.
//Example:
//for N = 4 and S = {1,2,3}, there are four solutions: {1,1,1,1},{1,1,2},{2,2},{1,3}.
//So output should be 4.
func CoinChangeProblem(coins []int, sum int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, sum+1)
	}

	for coinIndex := 1; coinIndex <= len(coins); coinIndex++ {
		for sumIndex := 1; sumIndex <= sum; sumIndex++ {
			coin := coins[coinIndex-1]
			if coin > sumIndex {
				notTaken := dp[coinIndex-1][sumIndex]
				dp[coinIndex][sumIndex] = notTaken
				continue
			}
			taken := coin + dp[coinIndex][sumIndex-coin]
			notTaken := dp[coinIndex-1][sumIndex]
			if taken > notTaken {
				dp[coinIndex][sumIndex] = taken
			} else {
				dp[coinIndex][sumIndex] = notTaken
			}
		}
	}

	return dp[len(coins)][sum]
}

//Coin Change Problem Minimum Numbers of coins
//Given a value V, if we want to make change for V cents, and we have infinite
//supply of each of C = { C1, C2, .. , Cm} valued coins, what is the minimum number of coins to make the change?
//Example:
//Input: coins[] = {25, 10, 5}, V = 30
//Output: Minimum 2 coins required
//We can use one coin of 25 cents and one of 5 cents
func MinNumberOfCoins(coins []int, sum int) {

}
