package dunzo

//Three Palindromic Substrings
//Given an input string word, split the string into exactly three non-empty palindromic substrings
//and return them. A string is called a palindrome if it reads the same forward and backward.
//	For example, abcba, a, and abba are palindromes, while abab and cd are not.
//	A substring of another string can be obtained from the original string by dropping 0 or more characters
//from the beginning, the end, or both ends. For example, the substrings of abc = [abc, ab, bc, a, b, c].
//	A substring is a palindromic substring if the substring is a palindrome.
//Working from left to right, choose the smallest split for the first item that
//still allows the remaining word to be split into 2 palindromes. Similarly,
//choose the smallest second palindromic substring that leaves a third palindromic substring.
//	Return the substrings in that order. If there is no way to split the word into exactly
//three palindromic substrings, return an array that contains a single string: Impossible.
//Example
//word = "kayakracecartenet" This string it can be split into kayak,
//racecar and tenet. Note that the smallest split that still leaves two more palindromes is kayak.
//	and the next palindrome is racecar.
//	The return array is kayak", "racecar", "tenet".
//word = "aardvark" There is no way to split it into three palindromic

/*
 * Complete the 'threePalindromicSubstrings' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING word as parameter.
 */

func ThreePalindromicSubstrings(word string) []string {
	wordLen := len(word)
	sol := []string{"Impossible"}
	if wordLen < 4 {
		return sol
	}

outer:
	for i := 1; i < wordLen-1; i++ {
		first := word[0:i]

		if !IsPalindromic(first) {
			continue
		}
		for j := i + 1; j < wordLen; j++ {
			second := word[i:j]

			if !IsPalindromic(second) {
				continue
			}
			third := word[j:]

			if !IsPalindromic(third) {
				continue
			}
			sol = []string{first, second, third}
			break outer
		}
	}

	return sol
}

func IsPalindromic(value string) bool {

	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}

	return true
}

func ThreePalindromicSubstringsWithMemo(word string) []string {
	wordLen := len(word)
	sol := []string{"Impossible"}
	if wordLen < 4 {
		return sol
	}
	memo := make(ThreePalindromicSubstringsMemo)

outer:
	for i := 1; i < wordLen-1; i++ {
		first := word[0:i]
		if !memo.IsPalindromic(first, 0, i) {
			continue
		}
		for j := i + 1; j < wordLen; j++ {
			second := word[i:j]
			if !memo.IsPalindromic(second, i, j) {
				continue
			}

			third := word[j:]
			if !memo.IsPalindromic(third, j, wordLen) {
				continue
			}

			sol = []string{first, second, third}
			break outer
		}
	}

	return sol
}

type ThreePalindromicSubstringsMemo map[[2]int]bool

func (memo ThreePalindromicSubstringsMemo) IsPalindromic(value string, start, end int) bool {
	if v, ok := memo[[2]int{start, end}]; ok {
		return v
	}

	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			memo[[2]int{start, end}] = false
			return false
		}
	}
	for start >= end {
		memo[[2]int{start, end}] = true
		start++
		end--
	}

	return true
}
