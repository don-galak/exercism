package bowling

import "fmt"

const pinsPerFrame, framesPerGame, rollsPerFrame = 10, 10, 2

type Game struct{ framesCompleted, score, pinsUp, rolls, bonus1, bonus2 int }

func NewGame() *Game {
	return &Game{
		pinsUp: pinsPerFrame,
	}
}
func (g *Game) isComplete() bool {
	return g.framesCompleted >= framesPerGame && g.bonus1 == 0 && g.bonus2 == 0
}

func (g *Game) Roll(pins int) error {
	if err := g.validateRoll(pins); err != nil {
		return err
	}

	g.rolls++
	g.pinsUp -= pins
	multiplier := 1
	isFill := g.framesCompleted >= framesPerGame
	if isFill {
		multiplier = 0
	}
	multiplier += g.bonus1 + g.bonus2
	g.score += pins * multiplier
	g.bonus1 = g.bonus2
	g.bonus2 = 0
	if g.pinsUp == 0 || g.rolls == rollsPerFrame {
		if g.pinsUp == 0 && !isFill {
			if g.rolls == rollsPerFrame {
				// Spare.
				g.bonus1++
			} else {
				// Strike.
				g.bonus2++
			}
		}
		g.rolls = 0
		g.pinsUp = pinsPerFrame
		g.framesCompleted++
	}
	return nil
}

func (g *Game) Score() (int, error) {
	if !g.isComplete() {
		return 0, fmt.Errorf("Game incomplete: On frame %d", g.framesCompleted+1)
	}
	return g.score, nil
}

const (
	errNegative        = "can't roll negative pins %d"
	errInputPinsGTPins = "there are only %d pins up, can't roll %d"
	errGameComplete    = "game complete, can't roll %d"
)

func (g *Game) validateRoll(pins int) error {
	switch {
	case pins < 0:
		return fmt.Errorf(errNegative, pins)
	case pins > g.pinsUp:
		return fmt.Errorf(errInputPinsGTPins, g.pinsUp, pins)
	case g.isComplete():
		return fmt.Errorf(errGameComplete, pins)
	}
	return nil
}
