package allergies

var allergiesNames = [8]string{"cats", "pollen", "chocolate", "tomatoes", "strawberries", "shellfish", "peanuts", "eggs"}

const maxAllergies = 256

func Allergies(allergies uint) (out []string) {
	var allergyVal uint = 128

	if allergies > maxAllergies {
		allergies = allergies % maxAllergies
	}

	for i := 0; allergies > 0; i++ {
		if allergies >= allergyVal {
			allergies -= allergyVal
			out = append(out, allergiesNames[i])
		}
		allergyVal /= 2
	}

	return
}

func AllergicTo(allergies uint, allergen string) bool {
	allergens := Allergies(allergies)

	for _, a := range allergens {
		if allergen == a {
			return true
		}
	}

	return false
}
