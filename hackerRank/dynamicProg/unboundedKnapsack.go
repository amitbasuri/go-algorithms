package dynamicProg

import (
	"encoding/json"
	"github.com/kr/pretty"
)

type unboundedKnapsackProbSoln struct {
	Max int
	N   map[string]struct{}
}

func UnboundedKnapsack(k int, arr []int) int {

	solns := unboundedKnapsackProbSoln{
		Max: k,
		N:   map[string]struct{}{},
	}
	soln := map[int]int{}
	newRemaining := map[int]int{}
	for _, a := range arr {
		newRemaining[a] = 0
	}
	solns.recursiveUnboundedKnapsack(newRemaining, soln)
	pretty.Println(solns)
	return len(solns.N)

}

func (p *unboundedKnapsackProbSoln) recursiveUnboundedKnapsack(remaining map[int]int, taken map[int]int) {

	//base cond
	if sumIntMap(taken) == p.Max {
		key, _ := json.Marshal(taken)
		p.N[string(key)] = struct{}{}
		return
	}

	// condition
	for a, _ := range remaining {
		// take
		if sumIntMap(taken)+a <= p.Max {
			newTaken := map[int]int{}
			for index, element := range taken {
				newTaken[index] = element
			}
			newTaken[a]++
			p.recursiveUnboundedKnapsack(remaining, newTaken)

		} else {

			// never take
			newRemaining := map[int]int{}
			for index, element := range remaining {
				newRemaining[index] = element
			}

			delete(newRemaining, a)
			p.recursiveUnboundedKnapsack(newRemaining, taken)
		}
	}
}

func min(arr []int) int {
	s := arr[0]
	for _, a := range arr {
		if s <= a {
			s = a
		}
	}
	return s
}

func sumIntMap(arr map[int]int) int {
	var s int
	for k, v := range arr {
		s += k * v
	}
	return s
}
