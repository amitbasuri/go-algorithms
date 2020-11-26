package others

func opposite(n int) int {
	return 0 - n
}

func mod(n int) int {
	if n < 0 {
		return 0 - n
	}
	return n
}

func SolutionMaxNumber(A []int) int {

	modMap := make(map[int]int, 0)
	max := 0

	for _, n := range A {
		if seenModValue, ok := modMap[mod(n)]; ok {
			if seenModValue+n == 0 {
				if seenModValue > max {
					max = seenModValue
				}
			}
		} else {

			modMap[mod(n)] = n
		}
	}

	return max

}
