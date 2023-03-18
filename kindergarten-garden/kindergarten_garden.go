package kindergarten

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

var errInvalidDiagram = errors.New("invalid diagram")
var errOddNumberOfCups = errors.New("odd number of cups")
var errInvalidCupCodes = errors.New("invalid cup codes")
var errDuplicateChildren = errors.New("duplicate children")

var plantMap = map[string]string{"V": "violets", "R": "radishes", "G": "grass", "C": "clover"}
var plantReg = regexp.MustCompile(`V|R|G|C`)

type Garden map[string][]string

func (g *Garden) Plants(child string) ([]string, bool) {
	c, exists := (*g)[child]
	return c, exists
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	firstRow, secondRow, err := getPlantsRows(diagram)
	if err != nil {
		return nil, err
	}

	garden := make(Garden, len(children))

	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	j := 0
	for _, child := range sortedChildren {
		if _, exists := garden[child]; exists {
			return nil, errDuplicateChildren
		}
		garden[child] = []string{
			plantMap[string(firstRow[j])],
			plantMap[string(firstRow[j+1])],
			plantMap[string(secondRow[j])],
			plantMap[string(secondRow[j+1])]}

		j += 2
	}

	return &garden, nil
}

func getPlantsRows(diagram string) (firstRow string, secondRow string, err error) {
	plants := strings.Split(diagram, "\n")

	if len(plants) != 3 || len(plants[1]) != len(plants[2]) {
		return "", "", errInvalidDiagram
	}

	if !plantReg.MatchString(plants[1]) || !plantReg.MatchString(plants[2]) {
		return "", "", errInvalidCupCodes
	}

	if len(plants[1])%2 != 0 || len(plants[2])%2 != 0 {
		return "", "", errOddNumberOfCups
	}

	return plants[1], plants[2], nil
}
