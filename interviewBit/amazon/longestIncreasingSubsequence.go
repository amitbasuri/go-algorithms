package amazon

//Input : A = [1, 2, 1, 5]
//Output : 3
//Explanation : The sequence : [1, 2, 5]

func lis(A []int) int {
	return longestIncreasingSubsequenceLength(A)
}

func longestIncreasingSubsequenceLength(arr []int) int {

	incSubsequenceArr := make([]int, len(arr)+1)
	//max := 0
	//
	//for i := 0; i < len(arr); i++ {
	//	lcsAtIndex := 0
	//	for j := 0; i < i; j++ {
	//
	//		if arr[i-1] > max {
	//			incSubsequenceArr[i] = 1 + incSubsequenceArr[i-1]
	//			max = arr[i-1]
	//			continue
	//		}
	//		incSubsequenceArr[i] = incSubsequenceArr[i-1]
	//	}
	//}
	return incSubsequenceArr[len(arr)]
}

const MaxInt32 int32 = 1<<31 - 1
const MaxInt64 int64 = 1<<63 - 1
const MaxUint64 uint64 = 1<<64 - 1
const MaxInt int = 1<<63 - 1
const MinInt int = -1 << 63

func maxInt(arr ...int) int {
	max := arr[0]
	for _, a := range arr {
		if a > max {
			max = a
		}
	}
	return max
}

func minInt(arr ...int) int {
	min := arr[0]
	for _, a := range arr {
		if a < min {
			min = a
		}
	}
	return min
}
