package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var reg = regexp.MustCompile(`[^0-9]`)
var err = errors.New("")

func Number(phoneNumber string) (string, error) {
	phoneNumber = reg.ReplaceAllString(phoneNumber, "")
	l := len(phoneNumber)

	if l < 10 || l > 11 {
		return "", err
	}

	first := phoneNumber[:1]

	if l == 11 {
		if first != "1" {
			return "", err
		}

		if areaCode := phoneNumber[1:2]; areaCode == "0" || areaCode == "1" {
			return "", err
		}

		if exchangeCode := phoneNumber[4:5]; exchangeCode == "0" || exchangeCode == "1" {
			return "", err
		}

		return phoneNumber[1:], nil
	}

	exchangeCode := phoneNumber[3:4]
	if first == "1" || first == "0" || exchangeCode == "0" || exchangeCode == "1" {
		return "", err
	}

	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return num[:3], nil
}

func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}
	num = fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:])
	return num, nil
}
