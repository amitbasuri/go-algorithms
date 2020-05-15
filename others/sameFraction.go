package others

import (
	"fmt"
	"github.com/kr/pretty"
)

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
