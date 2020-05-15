package others



func SolutionBulb(A []int) int {
	//isShining := make(map[int]struct{}, 0)
	var shiningMoments int

	isTurnedOnShining := make(map[int]struct{}, 0)
	isTurnedOnNotShining := make(map[int]struct{}, 0)


	for _, bulbNumber := range A {
		isTurnedOnNotShining[bulbNumber]= struct{}{}
		if  bulbNumber == 1 {
			isTurnedOnShining[bulbNumber]= struct{}{}
			delete(isTurnedOnNotShining, bulbNumber)
		}
		if _, ok := isTurnedOnShining[bulbNumber-1]  ; ok {
			isTurnedOnShining[bulbNumber]= struct{}{}
			delete(isTurnedOnNotShining, bulbNumber)
		}

		furtherBulbNumber := bulbNumber + 1

		for {
			if _, ok := isTurnedOnNotShining[furtherBulbNumber]  ; !ok {
				break
			}

			isTurnedOnShining[furtherBulbNumber]= struct{}{}
			delete(isTurnedOnNotShining, furtherBulbNumber)
			furtherBulbNumber++

		}

		if len(isTurnedOnNotShining) == 0 {
			shiningMoments++
		}
	}

	return shiningMoments

}
