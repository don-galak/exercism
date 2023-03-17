package kindergarten

import (
	"errors"
	"strings"
)

var plantMap = map[string]string{"V": "violets", "R": "radishes", "G": "grass", "C": "clover"}

type Garden struct {
	children map[string][]string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//	diagram := `
//	VVCCGG
//	VVCCGG`

var errInvalidDiagram = errors.New("invalid diagram")
var errOddNumberOfCups = errors.New("odd number of cups")

func getPlants(diagram string) (firstRow string, secondRow string, err error) {
	plants := strings.Split(diagram, "\n")
	// first := plants[1]
	// second := plants[2]
	if len(plants) != 3 || len(plants[1]) != len(plants[2]) {
		return "", "", errInvalidDiagram
	}

	if len(plants[1])%2 != 0 || len(plants[2])%2 != 0 {
		return "", "", errOddNumberOfCups
	}

	return plants[1], plants[2], nil
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	firstRow, secondRow, err := getPlants(diagram)
	if err != nil {
		return nil, err
	}

	println(children[0], firstRow, secondRow)

	G := Garden{children: make(map[string][]string, len(children))}

	j := 0
	for _, child := range children {
		G.children[child] = []string{
			plantMap[string(firstRow[j])],
			plantMap[string(firstRow[j+1])],
			plantMap[string(secondRow[j])],
			plantMap[string(secondRow[j+1])]}

		j += 2
	}

	return &G, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	return g.children[child], true
}
