package arrays

//First Missing Integer
//Asked in: Model N InMobi Amazon
//Given an unsorted integer array, find the first missing positive integer.
//Example:
//Given [1,2,0] return 3,
//[3,4,-1,1] return 2,
//[-8, -7, -6] returns 1
//Your algorithm should run in O(n) time and use constant space.

func firstMissingPositive(A []int) int {
	sol := len(A)
	for i := 0; i < len(A); i++ {
		if A[i] <= 0 {
			A[i] = 0
		}
	}
	temp := 0
	for i := 0; i < len(A); i++ {
		if A[i] != 0 {
			temp = A[i]
			if A[i] < 0 {
				temp *= -1
			}
			if temp < len(A) {
				if A[temp] > 0 {
					A[temp] *= -1
				}
				if A[temp] == 0 {
					A[temp] = -1*len(A) + 1
				}
			}
			if temp == len(A) {
				sol = len(A) + 1
			}
		}
	}

	for i := 1; i < len(A); i++ {
		if A[i] >= 0 {
			sol = i
			break
		}
	}
	return sol
}
