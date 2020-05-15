package dynamicProg

type KnapSack struct {
	//MaxWeight          int
	//AvailableMaterials map[string]int
	//TakenMaterials     map[string]struct{}
}

func (k *KnapSack) WeightNew(remainingWeight int, availableMaterials []string,
	weights []int, stolenWeight int) int {
	if remainingWeight == 0 || len(availableMaterials) == 0 {
		return stolenWeight
	}

	//material := availableMaterials[0]
	weight := weights[0]

	if weight > remainingWeight {
		newAvailableMaterials := make([]string, len(availableMaterials))
		copy(newAvailableMaterials, availableMaterials)
		newAvailableMaterials = newAvailableMaterials[1:]

		newWeights := make([]int, len(weights))
		copy(newWeights, weights)
		newWeights = newWeights[1:]

		// skip material
		return k.Weight(remainingWeight, newAvailableMaterials, newWeights, stolenWeight)

	} else {
		//take
		newAvailableMaterials := make([]string, len(availableMaterials))
		copy(newAvailableMaterials, availableMaterials)
		newAvailableMaterials = newAvailableMaterials[1:]

		newWeights := make([]int, len(weights))
		copy(newWeights, weights)
		newWeights = newWeights[1:]

		taken := k.Weight(remainingWeight-weight,
			newAvailableMaterials, newWeights, stolenWeight+weight)

		// dont take
		notTaken := k.Weight(remainingWeight, newAvailableMaterials, newWeights, stolenWeight)

		if taken > notTaken {
			return taken
		} else {
			return notTaken
		}

	}

}

func (k *KnapSack) Weight(remainingWeight int, availableMaterials []string,
	weights []int, stolenWeight int) int {
	if remainingWeight == 0 || len(availableMaterials) == 0 {
		return stolenWeight
	}

	//material := availableMaterials[0]
	weight := weights[0]

	if weight > remainingWeight {
		newAvailableMaterials := make([]string, len(availableMaterials))
		copy(newAvailableMaterials, availableMaterials)
		newAvailableMaterials = newAvailableMaterials[1:]

		newWeights := make([]int, len(weights))
		copy(newWeights, weights)
		newWeights = newWeights[1:]

		// skip material
		return k.Weight(remainingWeight, newAvailableMaterials, newWeights, stolenWeight)

	} else {
		//take
		newAvailableMaterials := make([]string, len(availableMaterials))
		copy(newAvailableMaterials, availableMaterials)
		newAvailableMaterials = newAvailableMaterials[1:]

		newWeights := make([]int, len(weights))
		copy(newWeights, weights)
		newWeights = newWeights[1:]

		taken := k.Weight(remainingWeight-weight,
			newAvailableMaterials, newWeights, stolenWeight+weight)

		// dont take
		notTaken := k.Weight(remainingWeight, newAvailableMaterials, newWeights, stolenWeight)

		if taken > notTaken {
			return taken
		} else {
			return notTaken
		}

	}

}
