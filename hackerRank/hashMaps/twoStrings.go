package hashMaps

func twoStrings(s1 string, s2 string) string {
	s1Map := make(map[rune]struct{})
	//s2Map := make(map[rune]struct{})

	for _, r := range s1 {
		s1Map[r] = struct{}{}
	}

	for _, r := range s2 {
		//s2Map[r] = struct{}{}
		if _, ok := s1Map[r]; ok {
			return "YES"
		}
	}

	return "NO"

}
