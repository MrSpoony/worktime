package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := flag.String("s", "08:00", "start time")
	end := flag.String("e", "17:00", "end time")
	pause := flag.String("p", "00:30", "pause")
	flag.Parse()

	starttime, isValid := validateHoursAndMinutes(*start, false)
	if !isValid {
		fmt.Println("Starting time is malformed")
	}
	endtime, isValid := validateHoursAndMinutes(*end, false)
	if !isValid {
		fmt.Println("Starting time is malformed")
	}
	pausetime, isValid := validateHoursAndMinutes(*pause, true)
	if !isValid {
		fmt.Println("Starting time is malformed")
	}

	timestart := createTimeDifference(starttime)
	timeend := createTimeDifference(endtime)
	timepause := createTimeDifference(pausetime)

	fmt.Println(timestart, timeend, timepause)

	diff := calculateTimeDifference(timestart, timeend, timepause)
	hours := diff.Hours()
	fmt.Println(diff)
	fmt.Printf("%.2f\n", hours)
}

func validateHoursAndMinutes(input string, minutesAsStandard bool) ([2]int, bool) {
	inputtime := strings.Split(input, ":")
	if len(inputtime) > 2 {
		return [2]int{0, 0}, false
	}
	if len(inputtime) == 1 {
		if len(inputtime[0]) > 2 {
			return validateHoursAndMinutes(input[:2]+":"+input[2:], minutesAsStandard)
		}
		stuff, err := strconv.Atoi(inputtime[0])
		if err != nil {
			return [2]int{0, 0}, false
		}
		if minutesAsStandard {
			if stuff >= 60 {
				return [2]int{0, 0}, false
			}
			return [2]int{0, stuff}, true
		}
		if stuff > 24 {
			return [2]int{0, 0}, false
		}
		return [2]int{stuff, 0}, true
	}
	hours, err := strconv.Atoi(inputtime[0])
	if err != nil {
		return [2]int{0, 0}, false
	}
	minutes, err := strconv.Atoi(inputtime[1])
	if err != nil {
		return [2]int{0, 0}, false
	}
	if minutes >= 60 {
		return [2]int{0, 0}, false
	}
	if hours > 24 {
		return [2]int{0, 0}, false
	}
	return [2]int{hours, minutes}, true
}

func calculateTimeDifference(start, end, pause time.Duration) time.Duration {
	now := time.Now()
	futuretime := now.Add(end).Add(-start).Add(-pause)
	diff := futuretime.Sub(now)
	return diff
}

func createTimeDifference(tm [2]int) time.Duration {
	now := time.Now()
	duration := now.Add(time.Hour*time.Duration(tm[0]) + time.Minute*time.Duration(tm[1])).Sub(now)
	return duration
}
