package others

//func main() {
//
//	arr := []int{
//		0 ,0 ,0 ,1, 0, 0,
//		//0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0,
//	}
//	println(jumpingOnClouds(arr))
//}

func jumpingOnClouds(clouds []int) int {
	var jumps int
	var index int

	for {
		if index < len(clouds)-2 {
			index = nextJumpIndex(clouds, index)
			jumps++
		} else {
			break
		}
	}

	if index == len(clouds) - 2 {
		jumps++
	}

	return jumps
}

func nextJumpIndex(clouds []int, currIndex int) int {
	if currIndex < len(clouds) - 2 {
		if clouds[currIndex+2] == 0 {
			return currIndex + 2
		} else {
			return currIndex + 1
		}
	} else {
			return currIndex + 1
	}

}
