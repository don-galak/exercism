package bottlesong

import (
	"fmt"
	"strings"
)

const bottleFallsText = "And if one green bottle should accidentally fall,"

var numMap = map[int]string{0: "no", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 10: "Ten"}

func getBottlesOnWall(num int) string {
	if number, exists := numMap[num]; exists {
		return fmt.Sprintf("%s green bottles hanging on the wall,", number)
	}
	return "One green bottle hanging on the wall,"
}

func getResultingBottlesOnWall(num int) string {
	if number, exists := numMap[num]; exists {
		return fmt.Sprintf("There'll be %s green bottles hanging on the wall.", strings.ToLower(number))
	}
	return "There'll be one green bottle hanging on the wall."
}

func Recite(startBottles, takeDown int) []string {
	song := make([]string, 0)

	for i := 0; i < takeDown; i++ {
		if i > 0 {
			song = append(song, "")
		}

		song = append(song, getBottlesOnWall(startBottles), getBottlesOnWall(startBottles), bottleFallsText)
		startBottles--
		song = append(song, getResultingBottlesOnWall(startBottles))
	}

	return song
}
