package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubUnsort(t *testing.T) {
	A := []int{1, 2, 3, 5, 6, 13, 15, 16, 17, 13, 13, 15, 17, 17, 17, 17, 17, 19, 19}
	assert.Equal(t, []int{6, 11}, subUnsort(A))

	A = []int{16, 15, 16, 20}
	assert.Equal(t, []int{0, 1}, subUnsort(A))

}
