package weeknumbers

import (
	"fmt"
	"time"
)

const Day = time.Hour * 24

func Hello() {
	fmt.Println("Hello, World!")
}

type Week struct {
	Start time.Time
	End   time.Time
}

func NewWeek(start, end time.Time) *Week {
	return &Week{
		Start: start,
		End:   end,
	}
}

func AllWeeks(year int) []*Week {

	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	weeks := make([]*Week, 0)

	// Handle week 0
	if start.Weekday() == time.Monday {
		weeks = append(weeks, nil)
	}

	for true {
		next := addDays(start, mondayOffset(start))
		weeks = append(weeks, NewWeek(start, addDays(next, -1)))

		if next.Year() != year {
			break
		}
		start = next
	}

	return weeks
}

func addDays(t time.Time, days int) time.Time {
	return t.Add(Day * time.Duration(days))
}

// mondayOffset returns the number of days you'd need to add to a date the next monday.
func mondayOffset(t time.Time) int {
	weekday := t.Weekday()

	// Scale t.Weekday() so that Monday is zero and Sunday is 6
	scaledWeekday := weekday - 1
	if scaledWeekday < 0 {
		scaledWeekday = 6
	}

	return 7 - int(scaledWeekday)
}
