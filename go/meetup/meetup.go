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

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) (day int) {
	var count int

	for i := 1; i < 32; i++ {
		date := fmt.Sprintf("%d/%d/%d 15:04:05", month, i, year)
		d, _ := time.Parse(layout, date)
		weekday := d.Weekday().String()

		if weekday == days[wDay] && d.Year() == year {
			count++

			if wSched < Last && int(wSched) == count-1 {
				return i
			}

			if wSched == Teenth && i > 12 && i < 20 {
				return i
			}
			day = i
		}
	}
	return
}
