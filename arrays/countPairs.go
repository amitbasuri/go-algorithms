package arrays

//Number of pairs
//Given two arrays X and Y of positive integers,
//find number of pairs such that xy > yx (raised to power of)
//where x is an element from X and y is an element from Y.
func countPairs(arr1 []int, arr2 []int) int {
	var count int
	for _, i := range arr1 {
		for _, j := range arr2 {
			if i^j > j^i {
				count++
			}
		}
	}

	return count
}
