package main

import (
	"time"
)

const Day = time.Hour * 24

type Week struct {
	Number int
	Start  time.Time
	End    time.Time
}

func NewWeek(number int, start, end time.Time) *Week {
	return &Week{
		Number: number,
		Start:  start,
		End:    end,
	}
}

func AllWeeks(year int) []*Week {

	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	weeks := make([]*Week, 0)
	weekNumber := 0

	// Not all years have a week 0
	if start.Weekday() == time.Monday {
		weekNumber++
	}

	for true {
		next := addDays(start, mondayOffset(start))
		end := addDays(next, -1)

		week := NewWeek(weekNumber, start, end)
		weeks = append(weeks, week)

		if next.Year() != year {
			break
		}
		start = next
		weekNumber++
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
