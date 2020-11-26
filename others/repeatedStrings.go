package others

//func main() {
//	println(repeatedString("a",1000000000000 ))
//
//}

func repeatedString(s string, n int64) int64 {
	sLen := int64(len(s))

	var sol int64
	rem := n % sLen
	qou := n / sLen
	var n1 int64
	var r2 int64

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "a" {
			n1++
			if int64(i) < rem {
				r2++
			}
		}

	}

	sol = qou*n1 + r2
	return sol

}
