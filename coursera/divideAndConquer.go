package coursera

import (
	"strconv"
)

// *************** week1 *******************
// recursive integer multiplication and/or Karatsuba's algorithm.
func multiply(m string, n string) string {
	if len(m) <= 2 && len(n) <= 2 {
		i, _ := strconv.Atoi(m)
		j, _ := strconv.Atoi(n)
		k := i * j
		return strconv.Itoa(k)
	}

	l1 := len(m) / 2
	a := m[:l1]
	b := m[l1:]

	l2 := len(n) / 2
	c := n[:l2]
	d := n[l2:]

	ac := multiply(a, c)
	ad := multiply(a, d)

	bc := multiply(b, c)
	bd := multiply(b, d)

	first := ac

	for z := 0; z < len(m)-l1; z++ {
		first += "0"
	}

	for z := 0; z < len(n)-l2; z++ {
		first += "0"
	}

	second := ad
	for z := 0; z < len(m)-l1; z++ {
		second += "0"
	}

	third := bc
	for z := 0; z < len(n)-l2; z++ {
		third += "0"
	}

	return sum(first, second, third, bd)

}

func sum(arrays ...string) string {
	leftOver := 0
	index := 1
	ans := ""
	for {
		found := false
		s := leftOver
		for _, arr := range arrays {
			if len(arr)-index < 0 {
				continue
			}
			found = true
			m := arr[len(arr)-index]
			n, _ := strconv.Atoi(string(m))
			s += n
		}
		index++
		if !found {
			if s != 0 {
				ans = strconv.Itoa(s) + ans
			}
			break
		} else {
			leftOver = s / 10
			ans = strconv.Itoa(s%10) + ans
		}

	}

	return ans
}

// *************** week2 *******************
func countSplitInversions(arr []int) ([]int, int) {
	arrLen := len(arr)
	if arrLen == 1 {
		return arr, 0
	}

	if arrLen == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
			return arr, 1
		}
		return arr, 0
	}

	half := len(arr) / 2

	leftArr, leftCount := countSplitInversions(arr[:half])
	rightArr, rightCount := countSplitInversions(arr[half:])
	splitArr, splitCount := countSplitInv(leftArr, rightArr)

	return splitArr, leftCount + rightCount + splitCount

}

func countSplitInv(arr1 []int, arr2 []int) ([]int, int) {

	sol := make([]int, len(arr1)+len(arr2))

	i := 0
	var count int
	rightInserted := 0

	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			sol[i] = arr1[0]
			arr1 = arr1[1:]
			count += rightInserted
		} else {
			sol[i] = arr2[0]
			arr2 = arr2[1:]
			rightInserted++
		}
		i++
	}

	for j := 0; j < len(arr1); j++ {
		sol[i] = arr1[j]
		count += rightInserted
		i++
	}
	for j := 0; j < len(arr2); j++ {
		sol[i] = arr2[j]
		i++
	}
	return sol, count
}
