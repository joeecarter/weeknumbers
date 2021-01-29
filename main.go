package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var year int

func parseYear(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return errors.New("expected either zero or one args (year)")
	}

	if len(args) == 0 {
		year = time.Now().Year()
		return nil
	}

	year64, err := strconv.ParseInt(args[0], 10, 32)
	year = int(year64)
	return err
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "weeknumbers",
		Short: "Creates lists of week numbers for a year in different formats.",
	}

	jsonCmd := &cobra.Command{
		Use:   "json [year]",
		Short: "Outputs a json array of week numbers.",
		Args:  parseYear,
		RunE:  runJsonCmd,
	}

	icalCmd := &cobra.Command{
		Use:   "ical [year]",
		Short: "Outputs a json array of week numbers.",
		Args:  parseYear,
		Run:   runIcalCmd,
	}

	rootCmd.AddCommand(jsonCmd)
	rootCmd.AddCommand(icalCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runJsonCmd(cmd *cobra.Command, args []string) error {
	weeks := AllWeeks(year)
	b, err := json.MarshalIndent(weeks, "", "\t")

	if err == nil {
		fmt.Println(string(b))
	}
	return err
}

func runIcalCmd(cmd *cobra.Command, args []string) {
	weeks := AllWeeks(year)

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)

	for _, week := range weeks {
		event := cal.AddEvent(uuid.NewString())
		event.SetSummary(week.Name())
		event.SetAllDayStartAt(week.Start)
		event.SetAllDayEndAt(week.End.Add(1 * Day))
	}

	fmt.Println(cal.Serialize())
}
