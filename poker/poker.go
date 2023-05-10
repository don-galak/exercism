package poker

import (
	"errors"
	"strings"
)

var ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
var suites = []string{"♢", "♡", "♤", "♧"}

func BestHand(hands []string) (o []string, e error) {

	// if len(hands) == 1 {
	// 	return hands, nil
	// }

	for _, hand := range hands {
		if !isValid(hand) {
			return nil, errors.New("invalid hand")
		}
	}

	return hands, nil
}

func isValid(hand string) bool {
	cards := strings.Fields(hand)

	if len(cards) != 5 {
		return false
	}

	for _, card := range cards {
		cardLen := len(card)

		if cardLen < 4 || cardLen > 5 {
			return false
		}

		if !(isFragmentValid(card[:cardLen-3], ranks) && isFragmentValid(card[cardLen-3:], suites)) {
			return false
		}
	}

	return true
}

func isFragmentValid(f string, cases []string) bool {
	for _, c := range cases {
		if f == c {
			return true
		}
	}
	return false
}
