package kindergarten

import "strings"

var plantMap = map[string]string{"V": "violets", "R": "radishes", "G": "grass", "C": "clover"}

type Garden struct {
	children map[string][]string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	plants := strings.Split(diagram, "\n")
	println(len(plants))

	G := Garden{children: make(map[string][]string, len(children))}

	// for i, j := 0, 0; i < len(children); i++ {
	// 	G.children[children[i]] = []string{plantMap[string(plants[0][j])], plantMap[string(plants[0][j+1])], plantMap[string(plants[1][j])], plantMap[string(plants[1][j+1])]}
	// }

	return &G, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	return []string{}, false
}
