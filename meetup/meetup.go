package meetup

import (
	"fmt"
	"time"
)

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Fifth
	Last
	Teenth
)

const layout = "1/2/2006 15:04:05"

var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	dada := []int{}

	for i := 1; i < 32; i++ {
		date := fmt.Sprintf("%d/%d/%d 15:04:05", month, i, year)
		d, _ := time.Parse(layout, date)
		weekday := d.Weekday().String()

		if weekday == days[wDay] && d.Year() == year {
			dada = append(dada, i)
		}
	}

	if wSched < Last {
		return dada[wSched]
	}

	if wSched == Last {
		return dada[len(dada)-1]
	}

	var lol int
	for _, d := range dada {
		if d > 12 && d < 20 {
			lol = d
		}
	}
	return lol
}
