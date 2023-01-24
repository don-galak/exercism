package allergies

const (
	eggs         uint = 1
	peanuts      uint = 2
	shellfish    uint = 4
	strawberries uint = 8
	tomatoes     uint = 16
	chocolate    uint = 32
	pollen       uint = 64
	cats         uint = 128
	maxAllergies uint = 256
)

func Allergies(allergies uint) (out []string) {
	if allergies > maxAllergies {
		allergies = allergies % maxAllergies
	}

	for allergies > 0 {
		switch {
		case allergies >= cats:
			allergies -= cats
			out = append(out, "cats")
		case allergies >= pollen:
			allergies -= pollen
			out = append(out, "pollen")
		case allergies >= chocolate:
			allergies -= chocolate
			out = append(out, "chocolate")
		case allergies >= tomatoes:
			allergies -= tomatoes
			out = append(out, "tomatoes")
		case allergies >= strawberries:
			allergies -= strawberries
			out = append(out, "strawberries")
		case allergies >= shellfish:
			allergies -= shellfish
			out = append(out, "shellfish")
		case allergies >= peanuts:
			allergies -= peanuts
			out = append(out, "peanuts")
		case allergies >= eggs:
			allergies -= eggs
			out = append(out, "eggs")
		}
	}

	return
}

func AllergicTo(allergies uint, allergen string) bool {
	return false
}
