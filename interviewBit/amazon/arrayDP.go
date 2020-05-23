package amazon

import (
	"strconv"
)

//Ways to Decode
//A message containing letters from A-Z is being encoded to numbers using the following mapping:
//
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//
//Input 2:
//A = "12"
//Output 2: 2
//Explanation 2:
//Given encoded message "12", it could be decoded as "AB" (1, 2) or "L" (12).
//The number of ways decoding "12" is 2.

func numDecodings(arr string) int {
	if len(arr) == 0 {
		return 0
	}
	dp := make([]int, len(arr)+1)
	if arr[0] == 48 {
		return 0
	}

	if arr == "10" || arr == "20" {
		return 1
	}
	dp[1] = 1

	for i := 2; i <= len(arr); i++ {
		if arr[i-1] == 48 {
			choice := arr[i-2 : i]
			if num, _ := strconv.Atoi(choice); num == 10 || num == 20 {
				dp[i] = dp[i-2]
				continue
			} else {
				return 0
			}
		}
		choice := arr[i-2 : i]
		if num, _ := strconv.Atoi(choice); num > 0 && num <= 26 {
			dp[i] = dp[i-1] + 1
			continue
		}
		dp[i] = dp[i-1]
	}

	return dp[len(arr)]
}
