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

type entry struct {
	date        string
	description string
	change      string
}

var errInvalidDate = errors.New("invalid date")
var localeMap = map[string]entry{
	"nl-NL": {
		date:        "Datum",
		description: "Omschrijving",
		change:      "Verandering",
	},
	"en-US": {
		date:        "Date",
		description: "Description",
		change:      "Change",
	},
}

func createHeader(locale string) (header string, err error) {
	if entry, ok := localeMap[locale]; !ok {
		return "", errors.New("invalid locale")
	} else {
		desc := entry.description
		date := entry.date
		ch := entry.change

		return date +
			strings.Repeat(" ", 10-len(date)) +
			" | " +
			desc +
			strings.Repeat(" ", 25-len(desc)) +
			" | " + ch + "\n", nil
	}
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
