package codility

//DNA sequence
func NaiveDnaSeqSolution(S string, P []int, Q []int) []int {

	sol := make([]int, len(P), len(P))

	factors := map[rune]int{
		'A': 1, 'C': 2, 'G': 3, 'T': 4,
	}

	curr := rune(S[0])
	var i int
	for j, n := range S {
		if curr != n {
			break
		}
		i = j
	}

	if i == len(S)-1 {
		f := factors[curr]
		for index := 0; index < len(sol); index++ {
			sol[index] = f
		}
		return sol
	}

	for i := 0; i < len(P); i++ {

		p := P[i]
		q := Q[i] + 1
		smallest := 5
		for _, n := range S[p:q] {
			if smallest == 1 {
				break
			}
			factor := factors[n]
			if factor < smallest {
				smallest = factor
			}
		}

		sol[i] = smallest
	}
	return sol
}

func DnaSeqSolution(S string, P []int, Q []int) []int {

	sol := make([]int, len(P), len(P))

	prefixSum := make([][]int, 3, 3)
	for i := 0; i < 3; i++ {
		prefixSum[i] = make([]int, len(S)+1, len(S)+1)
	}

	// calculate prefix sum Arr for A,C,G
	for i := 1; i <= len(S); i++ {
		ch := S[i-1]
		var a, c, g int
		if ch == 'A' {
			a = 1
		}
		if ch == 'C' {
			c = 1
		}
		if ch == 'G' {
			g = 1
		}

		prefixSum[0][i] = prefixSum[0][i-1] + a
		prefixSum[1][i] = prefixSum[1][i-1] + c
		prefixSum[2][i] = prefixSum[2][i-1] + g
	}

	// Figure out whether A,C or G is present in between the indices or not
	for i := 0; i < len(P); i++ {
		end := Q[i] + 1
		beg := P[i]

		if prefixSum[0][end]-prefixSum[0][beg] > 0 {
			sol[i] = 1
			continue
		}
		if prefixSum[1][end]-prefixSum[1][beg] > 0 {
			sol[i] = 2
			continue
		}
		if prefixSum[2][end]-prefixSum[2][beg] > 0 {
			sol[i] = 3
			continue
		}
		sol[i] = 4
	}

	return sol

}

func MinAvgTwoSliceSolution(A []int) int {
	// min slice len will always be less than 4 ,
	// imagine a slice of len 4 which can be divided in two slices of
	//len 2 which either has same avg or
	//any one of them having lesser avg (the other one will be greater)

	var start int
	min := float64(A[1]+A[0]) / 2

	for i := 2; i < len(A); i++ {
		avg := float64(A[i-1]+A[i]) / 2
		if avg < min {
			min = avg
			start = i - 1
		}

		avg = float64(A[i-2]+A[i-1]+A[i]) / 3
		if avg < min {
			min = avg
			start = i - 2
		}
	}

	return start

}

func PassingCarsSolution(A []int) int {

	suffixSumEast := make([]int, len(A)+1)

	var sol int

	for i := len(A) - 1; i >= 0; i-- {

		if A[i] == 1 { //west
			suffixSumEast[i] = suffixSumEast[i+1] + 1
		} else {
			suffixSumEast[i] = suffixSumEast[i+1]
			sol += suffixSumEast[i]
		}

		if sol > 1000000000 {
			return -1
		}

	}

	return sol

}
