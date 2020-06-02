package dunzo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRichCoinAndEnergy(t *testing.T) {
	assert.Equal(t, int32(12), GetRichCoinAndEnergy(0, []int32{2, 1, 1}, []int32{11, 5, 7}))
	assert.Equal(t, int32(18), GetRichCoinAndEnergy(1, []int32{2, 1, 1}, []int32{11, 5, 7}))
	assert.Equal(t, int32(23), GetRichCoinAndEnergy(2, []int32{2, 1, 1}, []int32{11, 5, 7}))

	assert.Equal(t, int32(13), GetRichCoinAndEnergy(1, []int32{1, 1, 1}, []int32{6, 5, 7}))
	assert.Equal(t, int32(13), GetRichCoinAndEnergy(1, []int32{1, 1, 1}, []int32{5, 6, 7}))
	assert.Equal(t, int32(20), GetRichCoinAndEnergy(0, []int32{1, 1, 1}, []int32{20, 6, 7}))

}

func TestThreePalindromicSubstrings(t *testing.T) {
	assert.Equal(t, []string{"kayak", "racecar", "tenet"}, ThreePalindromicSubstrings("kayakracecartenet"))
	assert.Equal(t, []string{"kayak", "racecar", "tenet"}, ThreePalindromicSubstringsWithMemo("kayakracecartenet"))
}

//func main() {
//	l := 1 << 12
//	i := make([]rune, l)
//	for j := 0; j < l; j++ {
//		i[j] = 'h'
//	}
//	t := time.Now()
//	s := threePalindromicSubstrings(string(i))
//	fmt.Println(s, time.Since(t).Milliseconds())
//}
