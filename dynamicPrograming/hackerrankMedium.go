package dynamicPrograming

//https://www.hackerrank.com/challenges/equal/problem
//Equal
func equal(arr []int, size int, options []int, optionsIndex int) int {

	if size == 0 {
		return 0
	}

	if allValuesEqual(arr) {
		return 0
	}

	return 0

}
func allValuesEqual(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	first := arr[0]
	for _, v := range arr {
		if v != first {
			return false
		}
	}
	return true

}

func sum(array ...int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
