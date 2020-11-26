package hashmaps

func CountTriplets(arr []int64, r int64) int64 {
	var total int64

	countMap := make(map[int64][]int64, 0)

	for i, n := range arr {
		if _, ok := countMap[n]; !ok {
			countMap[n] = []int64{int64(i)}
		} else {
			countMap[n] = append(countMap[n], int64(i))
		}
	}

	for n, indices := range countMap {
		if r == 1 {
			num := int64(len(indices))
			if num < 3 {
				total += num
				continue
			}
			total += num * (num - 1) * (num - 2) / (3 * 2)
			continue
		}

		c1, ok := countMap[n*r]
		if !ok {
			continue
		}
		c2, ok := countMap[n*r*r]
		if !ok {
			continue
		}
		total += int64(len(indices) * len(c1) * len(c2))
	}

	return total
}
