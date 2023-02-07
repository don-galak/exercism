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

var headerMap = map[string]string{
	"nl-NL": "Datum      | Omschrijving              | Verandering\n",
	"en-US": "Date       | Description               | Change\n",
}

var currencyMap = map[string]string{
	"EUR": "€",
	"USD": "$",
}

func createHeader(locale string) (string, error) {
	if _, ok := headerMap[locale]; !ok {
		return "", errInvalidLocale
	}
	return headerMap[locale], nil
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
			change.WriteString(currencyMap[currency])

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
			change.WriteString(currencyMap[currency])

			for i := len(parts) - 1; i >= 0; i-- {
				change.WriteString(parts[i] + ",")
			}
			change.Truncate(change.Len() - 1)
			change.WriteString(".")
			change.WriteString(centsStr[len(centsStr)-2:])
			if negative {
				c := change.String()
				change.Reset()
				change.WriteString("(" + c + ")")
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
