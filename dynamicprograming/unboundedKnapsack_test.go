package dynamicprograming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRodCuttingProblem(t *testing.T) {
	assert.Equal(t, 22, RodCuttingProblem([]int{1, 2, 3, 4, 5, 6, 7, 8},
		[]int{1, 5, 8, 9, 10, 17, 17, 20}, 8, 8))

	assert.Equal(t, 22, RodCuttingProblemBU([]int{1, 2, 3, 4, 5, 6, 7, 8},
		[]int{1, 5, 8, 9, 10, 17, 17, 20}, 8, 8))
}

func TestCoinChangeProblem(t *testing.T) {
	assert.Equal(t, 4, CoinChangeProblem([]int{1, 2, 3}, 4))
	assert.Equal(t, 5, CoinChangeProblem([]int{1, 2, 3}, 5))

}
