package queenattack

import (
	"errors"
	"math"
)

var abs = math.Abs

func invalidReturn() (float64, float64, float64, float64, error) {
	return 0, 0, 0, 0, errors.New("invalid positions")
}

func validatePositions(white, black string) (float64, float64, float64, float64, error) {
	if len(white) == 0 || white == black {
		return invalidReturn()
	}

	whiteX := float64(white[0] - 'a')
	whiteY := float64(white[1] - '0' - 1)
	blackX := float64(black[0] - 'a')
	blackY := float64(black[1] - '0' - 1)

	if whiteX > 7 || whiteY > 7 || blackX > 7 || blackY > 7 {
		return invalidReturn()
	}

	return whiteX, whiteY, blackX, blackY, nil
}

func CanQueenAttack(white, black string) (bool, error) {
	whiteX, whiteY, blackX, blackY, err := validatePositions(white, black)

	if err != nil {
		return false, err
	}

	switch {
	case whiteX == blackX:
		return true, nil
	case whiteY == blackY:
		return true, nil
	case abs(whiteX-blackX) == abs(whiteY-blackY):
		return true, nil
	}

	return false, nil
}
