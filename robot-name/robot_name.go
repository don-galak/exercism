package robotname

import (
	"errors"
	"fmt"
	"strings"
)

var namesGenerated = 0

const maxAvailableNames = 676001

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		err := r.Reset()
		if err != nil {
			return "", err
		}
	}
	return r.name, nil
}
func (r *Robot) Reset() error {
	namesGenerated++
	if namesGenerated >= maxAvailableNames {
		return errors.New("namespace is exhausted")
	}
	r.name = generateAlpha()
	return nil
}
func generateAlpha() string {
	representationForFirstChar := namesGenerated / 1000
	var w strings.Builder
	w.WriteRune(rune('A' + (representationForFirstChar / 26)))
	representationForSecondChar := representationForFirstChar % 26
	if representationForSecondChar > 0 {
		w.WriteRune(rune('A' + rune(representationForSecondChar)))
	} else {
		w.WriteRune('A')
	}
	// populate with digits
	w.WriteString(fmt.Sprintf("%03d", namesGenerated%1000))
	return w.String()
}
