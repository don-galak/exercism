package romannumerals

import (
	"errors"
)

var unitMap = map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX"}
var tenMap = map[int]string{1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC"}
var hundredMap = map[int]string{1: "C", 2: "CC", 3: "CCC", 4: "CD", 5: "D", 6: "DC", 7: "DCC", 8: "DCCC", 9: "CM"}
var thousandMap = map[int]string{1: "M", 2: "MM", 3: "MMM"}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("invalid input")
	}
	thousands := 0
	hundreds := 0
	tens := 0
	units := 0
	romanNumeral := ""

	if input > 999 {
		thousands = input / 1000
		if thousands > 0 {
			romanNumeral += thousandMap[thousands]
		}
	}

	if input > 99 {
		hundreds = (input - thousands*1000) / 100
		if hundreds > 0 {
			romanNumeral += hundredMap[hundreds]
		}
	}

	if input > 9 {
		tens = (input - thousands*1000 - hundreds*100) / 10
		if tens > 0 {
			romanNumeral += tenMap[tens]
		}
	}

	units = input % 10
	if units > 0 {
		romanNumeral += unitMap[units]
	}

	return romanNumeral, nil
}
