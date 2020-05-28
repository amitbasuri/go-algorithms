package arrays

//Find Duplicate in Array
//Asked in: Amazon VMWare Riverbed
//
//Given a read only array of n + 1 integers between 1 and n,
//find one number that repeats in linear time using less than O(n) space
//and traversing the stream sequentially O(1) times.
func repeatedNumber(A []int) int {
	s := make(map[int]struct{})
	for _, a := range A {
		if _, ok := s[a]; ok {
			return a
		} else {
			s[a] = struct{}{}
		}
	}
	return -1
}
