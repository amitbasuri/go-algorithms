package dynamicPrograming

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubsequenceTD(t *testing.T) {
	assert.Equal(t, 4, LongestCommonSubsequenceTD("abcdgh", "abedfh", 6, 6))
	assert.Equal(t, 2, LongestCommonSubsequenceTD("abcd", "xadct", 4, 5))

}

func TestLongestCommonSubstringLength(t *testing.T) {
	assert.Equal(t, 2, LongestCommonSubstringLength("abcdgh", "abedfh", 6, 6))
	assert.Equal(t, 2, LongestCommonSubstringLength("abcd", "xadcdt", 4, 6))
	assert.Equal(t, 0, LongestCommonSubstringLength("qwer", "xaxac", 4, 5))
}
