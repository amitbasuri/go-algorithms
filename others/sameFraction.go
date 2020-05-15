package main

import (
	"fmt"
	"github.com/kr/pretty"
)

func main() {
	//println(Solution([]int{1, 2, 3, 4, 0}, []int{2, 3, 6, 8, 4}))

	println(Solution([]int{3,3,4}, []int{5,4,3}))

}

func qoutient(n int, d int) int {
	return n / d
}

func remainder(n int, d int) int {
	return n % d
}

func root(n int, d int) string {
	if n > d {
		return fmt.Sprintf("%d_%d", qoutient(n, d), remainder(n, d))
	}

	if n == 0 {
		return fmt.Sprintf("0_0")

	}
	return fmt.Sprintf("%d_%d", remainder(d, n), qoutient(d, n))

}

func Solution(X []int, Y []int) int {
	max := 1
	maxFractionMap := make(map[string]int, 0)
	arrLen := len(X)

	for i := 0; i < arrLen; i++ {

		key := root(X[i], Y[i])
		maxFractionMap[key]++
		if maxFractionMap[key] > max {
			max = maxFractionMap[key]
		}
	}
	pretty.Pr

	return max
}
