package ledger

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var errInvalidDate = errors.New("invalid date")

func createHeader(locale string) (header string, err error) {
	if locale == "nl-NL" {
		header = "Datum" +
			strings.Repeat(" ", 10-len("Datum")) +
			" | " +
			"Omschrijving" +
			strings.Repeat(" ", 25-len("Omschrijving")) +
			" | " + "Verandering" + "\n"
	} else if locale == "en-US" {
		header = "Date" +
			strings.Repeat(" ", 10-len("Date")) +
			" | " +
			"Description" +
			strings.Repeat(" ", 25-len("Description")) +
			" | " + "Change" + "\n"
	} else {
		return "", errors.New("invalid locale")
	}
	return
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	header, err := createHeader(locale)
	if err != nil {
		return "", err
	}
	if !(currency == "USD" || currency == "EUR") {
		return "", errors.New("invalid currency")
	}

	var entriesCopy []Entry
	entriesCopy = append(entriesCopy, entries...)

	sort.SliceStable(entriesCopy, func(i, j int) bool {
		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	for _, entry := range entriesCopy {
		if len(entry.Date) != 10 {
			return "", errInvalidDate
		}
		d1, d2, d3, d4, d5 := entry.Date[:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:]
		if d2 != '-' {
			return "", errInvalidDate
		}
		if d4 != '-' {
			return "", errInvalidDate
		}
		description := entry.Description
		if len(description) > 25 {
			description = description[:22] + "..."
		} else {
			description = description + strings.Repeat(" ", 25-len(description))
		}
		var d string
		if locale == "nl-NL" {
			d = d5 + "-" + d3 + "-" + d1
		} else if locale == "en-US" {
			d = d3 + "/" + d5 + "/" + d1
		}
		negative := false
		cents := entry.Change
		if cents < 0 {
			cents *= -1
			negative = true
		}
		var row string
		if locale == "nl-NL" {
			if currency == "EUR" {
				row += "€"
			} else if currency == "USD" {
				row += "$"
			}
			row += " "
			centsStr := strconv.Itoa(cents)
			switch len(centsStr) {
			case 1:
				centsStr = "00" + centsStr
			case 2:
				centsStr = "0" + centsStr
			}
			rest := centsStr[:len(centsStr)-2]
			var parts []string
			for len(rest) > 3 {
				parts = append(parts, rest[len(rest)-3:])
				rest = rest[:len(rest)-3]
			}
			if len(rest) > 0 {
				parts = append(parts, rest)
			}
			for i := len(parts) - 1; i >= 0; i-- {
				row += parts[i] + "."
			}
			row = row[:len(row)-1]
			row += ","
			row += centsStr[len(centsStr)-2:]
			if negative {
				row += "-"
			} else {
				row += " "
			}
		} else if locale == "en-US" {
			if negative {
				row += "("
			}
			if currency == "EUR" {
				row += "€"
			} else if currency == "USD" {
				row += "$"
			} else {
				return "", errInvalidDate
			}
			centsStr := strconv.Itoa(cents)
			switch len(centsStr) {
			case 1:
				centsStr = "00" + centsStr
			case 2:
				centsStr = "0" + centsStr
			}
			rest := centsStr[:len(centsStr)-2]
			var parts []string
			for len(rest) > 3 {
				parts = append(parts, rest[len(rest)-3:])
				rest = rest[:len(rest)-3]
			}
			if len(rest) > 0 {
				parts = append(parts, rest)
			}
			for i := len(parts) - 1; i >= 0; i-- {
				row += parts[i] + ","
			}
			row = row[:len(row)-1]
			row += "."
			row += centsStr[len(centsStr)-2:]
			if negative {
				row += ")"
			} else {
				row += " "
			}
		} else {
			return "", errInvalidDate
		}
		var al int
		for range row {
			al++
		}
		header += d + strings.Repeat(" ", 10-len(d)) + " | " + description + " | " +
			strings.Repeat(" ", 13-al) + row + "\n"
	}
	return header, nil
}
