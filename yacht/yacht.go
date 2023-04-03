package yacht

func Score(dice []int, category string) int {
	diceMap := make(map[int]int)

	for _, d := range dice {
		diceMap[d] = diceMap[d] + 1
	}

	switch category {
	case "yacht":
		if len(diceMap) == 1 {
			return 50
		}
	case "full house":
		if len(diceMap) == 2 {
			result := 0
			for k, v := range diceMap {
				if v > 3 {
					return 0
				}
				result += k * v
			}
			return result
		}
	case "four of a kind":
		for k, v := range diceMap {
			if v >= 4 {
				return k * 4
			}
		}
	case "little straight":
		sum := 0
		for k := range diceMap {
			sum += k
		}
		if sum == 15 {
			return 30
		}
	case "big straight":
		sum := 0
		for k := range diceMap {
			sum += k
		}
		if sum == 20 {
			return 30
		}
	case "choice":
		sum := 0
		for k, v := range diceMap {
			sum += k * v
		}
		return sum
	case "ones":
		return diceMap[1] * 1
	case "twos":
		return diceMap[2] * 2
	case "threes":
		return diceMap[3] * 3
	case "fours":
		return diceMap[4] * 4
	case "fives":
		return diceMap[5] * 5
	case "sixes":
		return diceMap[6] * 6
	}

	return 0
}
