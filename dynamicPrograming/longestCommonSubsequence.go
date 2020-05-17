package dynamicPrograming

func LongestCommonSubsequenceTD(s1 string, s2 string, len1 int, len2 int) int {

	if len1 == 0 || len2 == 0 {
		return 0
	}

	if s1[len1-1] == s2[len2-1] {
		return 1 + LongestCommonSubsequenceTD(s1, s2, len1-1, len2-1)
	}

	choice1 := LongestCommonSubsequenceTD(s1, s2, len1, len2-1)
	choice2 := LongestCommonSubsequenceTD(s1, s2, len1-1, len2)

	if choice1 > choice2 {
		return choice1
	}

	return choice2

}

func ShortestCommonSuperSubsequence(s1 string, s2 string, len1 int, len2 int) int {
	return len1 + len2 - LongestCommonSubsequenceTD(s1, s2, len1, len2)
}

func LongestCommonSubstringLength(s1 string, s2 string, len1 int, len2 int) int {

	buArr := make(LongestCommonSubstringLengthBottomUpArr, len1+1) //make arr size adding 1
	for i := 0; i <= len1; i++ {
		buArr[i] = make([]int, len2+1) //make arr size adding 1
	}

	i, j := buArr.Fill(s1, s2, len1, len2)
	return buArr[i][j]
}

type LongestCommonSubstringLengthBottomUpArr [][]int

func (buArr LongestCommonSubstringLengthBottomUpArr) Fill(s1 string, s2 string, len1 int, len2 int) (int, int) {
	var i, j int
	for s1Index := 1; s1Index <= len1; s1Index++ { // loop from 1 to len
		for s2Index := 1; s2Index <= len2; s2Index++ {
			if s1[s1Index-1] == s2[s2Index-1] { // index with -1
				buArr[s1Index][s2Index] = 1 + buArr[s1Index-1][s2Index-1]
				if buArr[s1Index][s2Index] > buArr[i][j] {
					i, j = s1Index, s2Index
				}
				continue
			}
			buArr[s1Index][s2Index] = 0

		}

	}

	return i, j

}

func MinNumInsertionDeletionToConvertS1ToS2(s1 string, s2 string, len1 int, len2 int) (int, int) {
	lcs := LongestCommonSubsequenceTD(s1, s2, len1, len2)

	if len1 > len2 {
		return len1 - lcs, len2 - lcs
	}
	return len2 - lcs, len1 - lcs

}

func LongestPalindromicSubsequenceLength(s string) int {
	l := len(s)
	s2Rune := make([]rune, l)
	for i, r := range s {
		s2Rune[l-1-i] = r
	}

	return LongestCommonSubsequenceTD(s, string(s2Rune), l, l)

}
