package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(weekSchedule WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	checkStart, checkEnd := 0, 0
	switch weekSchedule {
	case First:
		checkStart, checkEnd = 1, 7
	case Second:
		checkStart, checkEnd = 8, 14
	case Third:
		checkStart, checkEnd = 15, 21
	case Fourth:
		checkStart, checkEnd = 22, 28
	case Teenth:
		checkStart, checkEnd = 13, 19
	case Last:
		month += 1
		checkStart, checkEnd = -6, 0
	}

	time := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)
	for i := checkStart; i <= checkEnd; i++ {
		d := time.AddDate(0, 0, i)
		if d.Weekday() == weekday {
			return d.Day()
		}
	}
	return 0
}
