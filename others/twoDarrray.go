package others

import (
	"fmt"
)

func main() {
	minimumBribes([]int{1, 2, 5, 3, 7, 8, 6, 4})

}

func minimumBribes(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	var count int

	swapCount := make(map[int]int)

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				if _, ok := swapCount[items[i]]; ok {
					swapCount[items[i]]++
				} else {
					swapCount[items[i]] = 1
				}

				if _, ok := swapCount[items[i+1]]; ok {
					swapCount[items[i+1]]++
				} else {
					swapCount[items[i+1]] = 1
				}
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
				count++
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}

	var notMatched bool
	q := items
	for i := 0; i < len(q); i++ {
		if q[i] != i+1 || swapCount[i+1] > 2 {
			println(q[i], i+1, swapCount[i+1])
			notMatched = true
			break
		}
	}
	fmt.Printf("=%d", q)
	if notMatched {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(count)
	}
}

func swap(q []int, i int, j int) {
	q[i], q[j] = q[j], q[i]
}

func rotLeft(a []int32, d int32) []int32 {
	arrLen := int32(len(a))

	shift := d % arrLen

	l := a[:shift+1]
	r := a[shift:]

	return append(l, r...)

}

func hourglassSum(arr [][]int32) int32 {
	var inital bool
	var maxSum int32
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			sum := arr[i][j] +
				arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] +
				arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]

			if inital == false {
				maxSum = sum
				inital = true
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum

}
