package codility

func SolutionBulb(A []int) int {
	var moments int
	var sum int
	var expectedSum int

	for i, bulb := range A {
		expectedSum += i + 1
		sum += bulb

		if expectedSum == sum {
			moments++
		}
	}

	return moments

}

func SolutionPositiveNegative(A []int) int {
	negativeSet := make(map[int]struct{})
	positiveSet := make(map[int]struct{})
	zeroCount := 0

	var max int

	for _, num := range A {
		if num == 0 {
			zeroCount++
		}
		if num > 0 {
			positiveSet[num] = struct{}{}
			if _, ok := negativeSet[-num]; ok {
				if num > max {
					max = num
				}
			}
		}
		if num < 0 {
			negativeSet[num] = struct{}{}
			if _, ok := positiveSet[-num]; ok {
				if 0-num > max {
					max = 0 - num
				}
			}
		}

	}

	if zeroCount > max {
		max = zeroCount
	}
	return max

}

func SolutionNumeratorDenominator(nums []int, dens []int) int {

	if len(nums) == 0 {
		return 0
	}
	fractionCount := make(map[[2]int]int)
	max := 1

	for i := 0; i < len(nums); i++ {

		n := nums[i]
		d := dens[i]
		gcd := greatestCommonDivisor(n, d)
		fraction := [2]int{
			n / gcd, d / gcd,
		}

		if val, ok := fractionCount[fraction]; ok {
			newVal := val + 1
			fractionCount[fraction] = newVal
			if newVal > max {
				max = newVal
			}
		} else {
			fractionCount[fraction] = 1
		}
	}

	return max

}

func greatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}

	return greatestCommonDivisor(b, a%b)
}
