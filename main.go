package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

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
		Use:   "json",
		Short: "Outputs a json array of week numbers.",
		Args:  parseYear,
		RunE:  runJsonCmd,
	}

	rootCmd.AddCommand(jsonCmd)

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

// func main() {
//
//
// 	for number, week := range weeks {
// 		if week == nil {
// 			continue
// 		}
//
// 		fmt.Printf("Week %d - %s to %s\n", number, week.Start.Format(time.RFC3339Nano), week.End.Format(time.RFC3339Nano))
// 	}
//
// 	// b, _ := json.MarshalIndent(weeks, "", "\t")
// 	// fmt.Println(string(b))
// }
