package arrays

//Maximum Absolute Difference
//Asked in: Amazon
//You are given an array of N integers, A1, A2 ,…, AN. Return maximum value of f(i, j) for all 1 ≤ i, j ≤ N.
//f(i, j) is defined as |A[i] - A[j]| + |i - j|, where |x| denotes absolute value of x.
//
//For example,
//A=[1, 3, -1]
//f(1, 1) = f(2, 2) = f(3, 3) = 0
//f(1, 2) = f(2, 1) = |1 - 3| + |1 - 2| = 3
//f(1, 3) = f(3, 1) = |1 - (-1)| + |1 - 3| = 4
//f(2, 3) = f(3, 2) = |3 - (-1)| + |2 - 3| = 5
//
//So, we return 5.
func maxArr(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			v := abs(arr[i]-arr[j]) + abs(i-j)
			if v > max {
				max = v
			}
		}
	}
	return max
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
