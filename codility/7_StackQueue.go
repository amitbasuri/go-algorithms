package codility

// Fish
func FishSolution(sizes []int, directions []int) int {

	if len(sizes) == 0 {
		return 0
	}

	fishStack := make([]int, 0)
	fishStack = append(fishStack, 0)

	for currFishIndex := 1; currFishIndex < len(sizes); currFishIndex++ {

		fishStack = append(fishStack, currFishIndex) // 1-> 1-> , <-0 <-0 , 1-> <-0 , <-0 1->

		//more than 1 fish and top 2 fist opp dir
		for len(fishStack) > 1 {

			first := fishStack[len(fishStack)-1]
			second := fishStack[len(fishStack)-2]

			// 1-> 1-> , <-0 <-0
			if directions[first] == directions[second] {
				break
			}

			// <-0 1->
			if directions[second] == 0 && directions[first] == 1 {
				break
			}

			// 1-> <-0
			if sizes[second] > sizes[first] { // 2nd eats 1st
				fishStack = fishStack[:len(fishStack)-1] //pop
			} else { // 1st eats 2nd
				fishStack = fishStack[:len(fishStack)-1] //pop
				fishStack = fishStack[:len(fishStack)-1] //pop

				fishStack = append(fishStack, first) //add first
			}

		}

	}

	return len(fishStack)

}
