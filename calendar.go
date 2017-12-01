package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var usage = `
Usage:
	vimwiki-todo <MM> <YYYY>
`

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	startDate := calculateStart(os.Args[1:])
	fmt.Printf("# TODO: %s\n\n", startDate.Format("January 2006"))

	curDate := startDate
	for {
		switch {
		case int(curDate.Weekday()) == 0:
		case int(curDate.Weekday()) == 6:
		case int(curDate.Weekday()) == 5:
			fmt.Printf("== %s ==\n\n\n", curDate.Format("Mon, Jan 2"))
		default:
			fmt.Printf("== %s ==\n", curDate.Format("Mon, Jan 2"))
		}
		curDate = curDate.AddDate(0, 0, 1)
		if curDate.Month() != startDate.Month() {
			break
		}
	}

	fmt.Printf("\n\n== Unscheduled==\n\n==Notes==\n")
}

func calculateStart(args []string) time.Time {
	month, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "must be a numeric month:", args[0])
		os.Exit(2)
	}

	year, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "must be a valid year:", args[1])
	}

	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
}
