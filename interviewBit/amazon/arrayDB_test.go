package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	A := "875361268549483279131"

	assert.Equal(t, 6, numDecodings(A))
}
