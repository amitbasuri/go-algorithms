package dynamicPrograming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnapsackTD(t *testing.T) {
	assert.Equal(t, 220, KnapsackTD([]int{60, 100, 120}, []int{10, 20, 30}, 50, 3))
}

func TestKnapsackBU(t *testing.T) {
	assert.Equal(t, 220, KnapsackBU([]int{60, 100, 120}, []int{10, 20, 30}, 50, 3))
}
