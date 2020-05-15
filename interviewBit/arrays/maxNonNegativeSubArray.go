package arrays

//Input 1:
//A = [1, 2, 5, -7, 2, 3]
//
//Output 1:
//[1, 2, 5]

func MaxNonNegSubArr(inputArr []int) []int {
	inputArr = append(inputArr, -1)
	var startIndex, endIndex, sum int
	var maxStartIndex, maxEndIndex, maxSum int
	var startNew bool = true

	for i, n := range inputArr {
		if n < 0 {
			if startNew {
				continue
			}
			if i >= 0 {
				endIndex = i - 1

				if maxSum <= sum {
					maxStartIndex = startIndex
					maxEndIndex = endIndex
					maxSum = sum
				}
			}
			startNew = true
			continue
		}

		if startNew {
			startIndex = i
			startNew = false
			sum = 0
		}

		sum += n

	}

	if maxStartIndex == maxEndIndex && inputArr[maxStartIndex] < 0 {
		return []int{}
	}

	return inputArr[maxStartIndex : maxEndIndex+1]
}
