package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/MrSpoony/worktime/diff"
	"github.com/MrSpoony/worktime/validate"
)

func main() {
	start := flag.String("s", "08:00", "start time")
	end := flag.String("e", "17:00", "end time")
	breakstrb := flag.String("b", "", "break")
	breakstrp := flag.String("p", "00:30", "break")

	defaultgoal := flag.Bool("d", false, "reach default goal of 8.3")
	goal := flag.Float64("g", 0, "goal time, overrides default goal")

	flag.Parse()

	if *breakstrb != "" && *breakstrp != "" {
		fmt.Println("You can only use one break flag")
		return
	}

	breakstr := breakstrp
	if *breakstrp == "" {
		breakstr = breakstrb
	}

	var timebreak time.Duration
	if strings.Contains(*breakstr, "-") {
		parts := strings.Split(*breakstr, "-")
		if len(parts) != 2 {
			panic("Invalid break time, should either be a duration hh:mm or a range hh:mm-hh:mm")
		}
		startbreak, err := validate.HoursAndMinutes(parts[0], true)
		panicOnErr(err)
		endbreak, err := validate.HoursAndMinutes(parts[1], true)
		panicOnErr(err)
		startbreaktime := diff.IntsToTimeDiff(startbreak)
		endbreaktime := diff.IntsToTimeDiff(endbreak)
		timebreak = diff.GetTimeDiff(startbreaktime, endbreaktime, 0)
	} else {
		breakdur, err := validate.HoursAndMinutes(*breakstr, true)
		panicOnErr(err)
		timebreak = diff.IntsToTimeDiff(breakdur)
	}

	if *goal == 0 && *defaultgoal {
		*goal = 8.3
	}

	if *goal > 0 {
		timediff, err := validate.HoursAndMinutes(*start, false)
		panicOnErr(err)
		timestart := diff.IntsToTimeDiff(timediff)
		diff := time.Duration(float64(time.Hour) * *goal)
		timeend := timestart + timebreak + diff
		fmt.Printf("Start: %v, End: %v, Pause: %v\n", timestart, timeend, timebreak)
		fmt.Printf("Working time: %v\n", diff)
		fmt.Printf("%.2f\n", diff.Hours())
		return
	}

	starttime, err := validate.HoursAndMinutes(*start, false)
	panicOnErr(err)
	endtime, err := validate.HoursAndMinutes(*end, false)
	panicOnErr(err)

	timestart := diff.IntsToTimeDiff(starttime)
	timeend := diff.IntsToTimeDiff(endtime)

	fmt.Printf("Start: %v, End: %v, Pause: %v\n", timestart, timeend, timebreak)

	difference := diff.GetTimeDiff(timestart, timeend, timebreak)
	fmt.Printf("Working time: %v\n", difference)
	fmt.Printf("%.2f\n", difference.Hours())
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
