package arrays

func mergeSortedArrays(arr1 []int, arr2 []int) []int {

	sol := make([]int, len(arr1)+len(arr2))

	i := 0

	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			sol[i] = arr1[0]
			arr1 = arr1[1:]
		} else {
			sol[i] = arr2[0]
			arr2 = arr2[1:]
		}
		i++
	}

	for j := 0; j < len(arr1); j++ {
		sol[i] = arr1[j]
		i++
	}
	for j := 0; j < len(arr2); j++ {
		sol[i] = arr2[j]
		i++
	}
	return sol
}
