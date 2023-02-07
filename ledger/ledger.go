package ledger

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var errInvalidDate = errors.New("invalid date")
var errInvalidLocale = errors.New("invalid locale")
var errInvalidCurrency = errors.New("invalid currency")

type entry struct {
	date        string
	description string
	change      string
}

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
		return "", errInvalidLocale
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

const layout = "2006-01-02"
const nlLayout = "02-01-2006"
const usLayout = "01/02/2006"

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	header, err := createHeader(locale)
	if err != nil {
		return "", err
	}
	if !(currency == "USD" || currency == "EUR") {
		return "", errInvalidCurrency
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

		time, err := time.Parse(layout, entry.Date)
		if err != nil {
			return "", errInvalidDate
		}
		var date = formatDateBasedOnLocale(time, locale)

		description := entry.Description
		if len(description) > 25 {
			description = description[:22] + "..."
		} else {
			description = description + strings.Repeat(" ", 25-len(description))
		}
		negative := false
		cents := entry.Change
		if cents < 0 {
			cents *= -1
			negative = true
		}
		var change bytes.Buffer
		parts, centsStr := getPartsCents(cents)
		if locale == "nl-NL" {
			if currency == "EUR" {
				change.WriteString("€")
			} else if currency == "USD" {
				change.WriteString("$")
			}
			change.WriteString(" ")
			for i := len(parts) - 1; i >= 0; i-- {
				change.WriteString(parts[i] + ".")
			}
			change.Truncate(change.Len() - 1)
			change.WriteString(",")
			change.WriteString(centsStr[len(centsStr)-2:])
			if negative {
				change.WriteString("-")
			} else {
				change.WriteString(" ")
			}
		} else if locale == "en-US" {
			if negative {
				change.WriteString("(")
			}
			if currency == "EUR" {
				change.WriteString("€")
			} else if currency == "USD" {
				change.WriteString("$")
			}
			for i := len(parts) - 1; i >= 0; i-- {
				change.WriteString(parts[i] + ",")
			}
			change.Truncate(change.Len() - 1)
			change.WriteString(".")
			change.WriteString(centsStr[len(centsStr)-2:])
			if negative {
				change.WriteString(")")
			} else {
				change.WriteString(" ")
			}
		}

		spaces := 0
		for range change.String() {
			spaces++
		}
		spaces = 13 - spaces
		// panic(fmt.Sprint("\nrowLEN: ", change.Len(), "\nAL: ", spaces))

		header += date + strings.Repeat(" ", 10-len(date)) + " | " + description + " | " +
			strings.Repeat(" ", spaces) + change.String() + "\n"
	}
	return header, nil
}

func formatDateBasedOnLocale(time time.Time, locale string) string {
	switch locale {
	case "nl-NL":
		return time.Format(nlLayout)
	case "en-US":
		return time.Format(usLayout)
	default:
		return ""
	}
}

func getPartsCents(cents int) ([]string, string) {
	centsStr := fmt.Sprintf("%03d", cents)

	rest := centsStr[:len(centsStr)-2]
	var parts []string
	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}
	if len(rest) > 0 {
		parts = append(parts, rest)
	}

	return parts, centsStr
}
