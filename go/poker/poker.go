package poker

import (
	"bytes"
	"errors"
	"sort"
)

type group struct{ val, count byte }
type pokerHand struct {
	source string
	groups []group
	rank   int
}
type handSlice []pokerHand

func (h handSlice) Len() int           { return len(h) }
func (h handSlice) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h handSlice) Less(i, j int) bool { return h[i].compare(h[j]) > 0 }

const (
	straightFlush = iota
	fourOfAKind
	fullHouse
	flush
	straight
	threeOfAKind
	twoPair
	pair
	highCard
)

var valMap = map[rune]byte{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, '0': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
var suitMap = map[rune]byte{'♧': 0, '♢': 1, '♡': 2, '♤': 3}

func newPokerHand(in string, out *pokerHand) error {
	// Load array of value counters, and set of present suits
	runeIn := []rune(in)
	var valueCounter [15]int
	var suits [4]byte
	start, cards := 0, 0
	for end := range runeIn {
		if end == len(runeIn)-1 {
			end++
		} else if runeIn[end] != ' ' && end < len(runeIn)-1 {
			continue
		}
		slice := runeIn[start:end]
		if len(slice) != 2 && len(slice) != 3 {
			return errors.New("invalid card")
		}
		if v, ok := valMap[slice[len(slice)-2]]; !ok {
			return errors.New("invalid face value")
		} else {
			valueCounter[v]++
			cards++
		}
		if v, ok := suitMap[slice[len(slice)-1]]; !ok {
			return errors.New("invalid suit")
		} else {
			suits[v] = 1
		}
		start = end + 1
	}
	if cards != 5 {
		return errors.New("hand must have five cards")
	}
	// Load sorted slice of value grps
	grps := make([]group, 0, len(valueCounter))
	for count := 5; count > 0; count-- {
		for i := len(valueCounter) - 1; i > 0; i-- {
			if valueCounter[i] == count {
				grps = append(grps, group{byte(i), byte(valueCounter[i])})
			}
		}
	}
	// Check groups to determine result
	out.source, out.groups = in, grps
	if len(grps) == 5 {
		result := highCard
		if grps[0].val == 14 && grps[1].val == 5 && grps[4].val == 2 {
			out.groups = append(out.groups[1:], out.groups[0])
			result = straight
		} else if grps[4].val == grps[0].val-4 {
			result = straight
		}
		if bytes.Count(suits[:], []byte{1}) == 1 {
			if result == straight {
				result = straightFlush
			} else {
				result = flush
			}
		}
		out.rank = result
	} else if grps[0].count == 4 {
		out.rank = fourOfAKind
	} else if grps[0].count == 3 && grps[1].count == 2 {
		out.rank = fullHouse
	} else if grps[0].count == 3 {
		out.rank = threeOfAKind
	} else if grps[0].count == 2 {
		out.rank = twoPair
	} else {
		out.rank = pair
	}
	return nil
}
func BestHand(hands []string) ([]string, error) {
	pokerHands := make(handSlice, len(hands))
	for i, h := range hands {
		if err := newPokerHand(h, &pokerHands[i]); err != nil {
			return nil, err
		}
	}
	sort.Sort(pokerHands)
	results := hands[:0]
	for i := range pokerHands {
		if pokerHands[i].compare(pokerHands[0]) == 0 {
			results = append(results, pokerHands[i].source)
		}
	}
	return results, nil
}
func (h pokerHand) compare(other pokerHand) int {
	if h.rank != other.rank {
		return other.rank - h.rank
	} else if len(h.groups) != len(other.groups) {
		return len(other.groups) - len(h.groups)
	}
	for i := range h.groups {
		if h.groups[i].val > other.groups[i].val || h.groups[i].count > other.groups[i].count {
			return 1
		} else if h.groups[i].val < other.groups[i].val || h.groups[i].count < other.groups[i].count {
			return -1
		}
	}
	return 0
}
