package kindergarten

import (
	"errors"
	"regexp"
	"strings"
)

var errInvalidDiagram = errors.New("invalid diagram")
var errOddNumberOfCups = errors.New("odd number of cups")
var errInvalidCupCodes = errors.New("invalid cup codes")
var errDuplicateChildren = errors.New("duplicate children")

var plantMap = map[string]string{"V": "violets", "R": "radishes", "G": "grass", "C": "clover"}
var plantReg = regexp.MustCompile(`V|R|G|C`)

type Garden struct {
	children map[string][]string
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

func NewGarden(diagram string, children []string) (*Garden, error) {
	firstRow, secondRow, err := getPlantsRows(diagram)
	if err != nil {
		return nil, err
	}

	G := Garden{children: make(map[string][]string, len(children))}
	// var ch []string
	// copy(ch, children)
	// sort.Strings(ch)

	j := 0
	for _, child := range children {
		if _, exists := G.children[child]; exists {
			return nil, errDuplicateChildren
		}

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
	c, exists := g.children[child]
	return c, exists
}
