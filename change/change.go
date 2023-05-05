package change

import "errors"

func Change(coins []int, target int) (out []int, err error) {
	if target < 0 {
		return nil, errors.New("target cannot be negative")
	}
	results := make([][]int, target+1)
	results[0] = []int{}
	for amt := range results {
		for _, coin := range coins {
			if amt-coin >= 0 &&
				results[amt-coin] != nil &&
				(results[amt] == nil || len(results[amt-coin])+1 < len(results[amt])) {
				results[amt] = append([]int{coin}, results[amt-coin]...)
			}
		}
	}
	if results[target] == nil {
		return nil, errors.New("no solution")
	}
	return results[target], nil
}
