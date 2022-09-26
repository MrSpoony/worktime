package main

import (
	"flag"
	"fmt"

	"github.com/MrSpoony/worktime/diff"
	"github.com/MrSpoony/worktime/validate"
)

func main() {
	start := flag.String("s", "08:00", "start time")
	end := flag.String("e", "17:00", "end time")
	breakstr := flag.String("b", "00:30", "break")
	flag.Parse()

	starttime, err := validate.HoursAndMinutes(*start, false)
	panicOnErr(err)
	endtime, err := validate.HoursAndMinutes(*end, false)
	panicOnErr(err)
	breakdur, err := validate.HoursAndMinutes(*breakstr, true)
	panicOnErr(err)

	timestart := diff.IntsToTimeDiff(starttime)
	timeend := diff.IntsToTimeDiff(endtime)
	timebreak := diff.IntsToTimeDiff(breakdur)

	fmt.Printf("Start: %v, End: %v, Pause: %v\n", timestart, timeend, timebreak)

	difference := diff.GetTimeDiff(timestart, timeend, timebreak)
	hours := difference.Hours()
	fmt.Printf("Working time: %v\n", difference)
	fmt.Printf("%.2f\n", hours)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
