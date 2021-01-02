package main

import (
	"fmt"
	"time"

	"github.com/joeecarter/weeknumbers"
)

func main() {
	weeks := weeknumbers.AllWeeks(2021)

	for number, week := range weeks {
		if week == nil {
			continue
		}

		fmt.Printf("Week %d - %s to %s\n", number, week.Start.Format(time.RFC3339Nano), week.End.Format(time.RFC3339Nano))
	}

	// b, _ := json.MarshalIndent(weeks, "", "\t")
	// fmt.Println(string(b))
}
