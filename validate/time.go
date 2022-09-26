package validate

// TODO: Support for seconds?

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errMalformedInput = errors.New("malformed input")
)

// HoursAndMinutes calculates the hours and minutes from the input string
// the minutesFirst boolean tells the function what to do if there
// is a two digit number to parse, if true it will be a minute if false it's a hour
// The output is a slice with length 2 and the format []int{hours, minutes}
// The slice is nil if the error is not nil
func HoursAndMinutes(input string, minutesFirst bool) (output []int, err error) {
	inputtimes := strings.Split(input, ":")
	hours := 0
	minutes := 0
	defer func() {
		if output == nil && err == nil {
			output = []int{hours, minutes}
		}
		output, err = checkTime(output, err)
	}()
	if len(inputtimes) > 2 {
		err = errMalformedInput
		return
	}
	if len(inputtimes) == 1 {
		inputtime := inputtimes[0]
		switch len(inputtime) {
		case 4:
			fallthrough
		case 3:
			output, err = HoursAndMinutes(input[:len(input)-2]+":"+input[len(input)-2:], minutesFirst)
			return
		case 2:
			fallthrough
		case 1:
			var stuff int
			stuff, err = strconv.Atoi(inputtime)
			if err != nil {
				return
			}
			if minutesFirst {
				minutes = stuff
				return
			}
			hours = stuff
			return
		default:
			err = errMalformedInput
			return
		}
	}
	hours, err = strconv.Atoi(inputtimes[0])
	if err != nil {
		return
	}
	minutes, err = strconv.Atoi(inputtimes[1])
	if err != nil {
		return
	}
	return
}

func checkTime(input []int, err error) ([]int, error) {
	if err != nil {
		return nil, errMalformedInput
	}
	hours := input[0]
	minutes := input[1]
	if minutes >= 60 {
		return nil, errMalformedInput
	}
	if hours > 24 {
		return nil, errMalformedInput
	}
	return input, nil
}
