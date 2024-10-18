package yacht

var categoryMap = map[string]int{"ones": 1, "twos": 2, "threes": 3, "fours": 4, "fives": 5, "sixes": 6}

func Score(dice []int, category string) int {
	// key of diceMap is dice and value is occurences of dice
	diceMap := make(map[int]int)
	sum := 0

	for _, d := range dice {
		diceMap[d] = diceMap[d] + 1
		sum += d
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
		if len(diceMap) == 5 && diceMap[6] == 0 {
			return 30
		}
	case "big straight":
		if len(diceMap) == 5 && diceMap[1] == 0 {
			return 30
		}
	case "choice":
		return sum
	default:
		return diceMap[categoryMap[category]] * categoryMap[category]
	}

	return 0
}
