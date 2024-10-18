package twobucket

import "errors"

type bucket struct{ name, size, level int }

var names = map[int]string{1: "one", 2: "two"}

func Solve(thisSize, otherSize, goal int, start string) (string, int, int, error) {
	if err := validateInput(thisSize, otherSize, goal, start); err != nil {
		return "", 0, 0, err
	}

	this, other := &bucket{1, thisSize, 0}, &bucket{2, otherSize, 0}
	if start == names[2] {
		this, other = other, this
	}
	moves := 0
	for ; this.level != goal && other.level != goal; moves++ {
		switch {
		case this.level == 0:
			this.level = this.size
		case other.size == goal:
			other.level = goal
		case other.level == other.size:
			other.level = 0
		default:
			pour := other.size - other.level
			if this.level < pour {
				pour = this.level
			}
			this.level -= pour
			other.level += pour
		}
	}
	if other.level == goal {
		this, other = other, this
	}
	return names[this.name], moves, other.level, nil
}

func validateInput(thisSize, otherSize, goal int, start string) error {
	if thisSize <= 0 || otherSize <= 0 || goal <= 0 ||
		(goal > thisSize && goal > otherSize) ||
		(start != names[1] && start != names[2]) ||
		goal%gcd(thisSize, otherSize) != 0 {
		return errors.New("invalid input or no solution")
	}
	return nil
}

func gcd(a, b int) int {
	for ; b != 0; a, b = b, a%b {
	}
	return a
}
