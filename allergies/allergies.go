package allergies

var allergiesNames = []string{
	"cats",
	"pollen",
	"chocolate",
	"tomatoes",
	"strawberries",
	"shellfish",
	"peanuts",
	"eggs",
}

const maxAllergies = 256

func Allergies(allergies uint) (out []string) {
	var maxAllergyValue uint = 128

	if allergies > maxAllergies {
		allergies = allergies % maxAllergies
	}

	i := 0

	for allergies > 0 {
		if allergies >= maxAllergyValue {
			allergies -= maxAllergyValue
			out = append(out, allergiesNames[i])
		}
		i++
		maxAllergyValue /= 2
	}

	return
}

func AllergicTo(allergies uint, allergen string) bool {
	return false
}
