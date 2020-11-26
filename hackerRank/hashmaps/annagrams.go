package hashmaps

func SherlockAndAnagrams(s string) int32 {
	subStingMapAlpha := make(map[[26]int]int, 0)

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			ch := s[i : j+1]
			subStingMapAlpha[signature(ch)]++
		}
	}

	var sol int

	for _, count := range subStingMapAlpha {
		sol += count * (count - 1) / 2
	}

	return int32(sol)
}

func signature(s string) [26]int {
	a := [26]int{}

	for _, ch := range s {
		a[ch-97]++
	}

	return a
}
