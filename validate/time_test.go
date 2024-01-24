package validate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHoursAndMinutes(t *testing.T) {
	type testcase struct {
		input        string
		minutesFirst bool
		want         []int
		wantErr      bool
	}
	testcases := []testcase{
		{"0800", true, []int{8, 0}, false},
		{"08:00", true, []int{8, 0}, false},
		{"0800", false, []int{8, 0}, false},
		{"08:00", false, []int{8, 0}, false},
		{"08", false, []int{8, 0}, false},
		{"8", false, []int{8, 0}, false},
		{"800", true, []int{8, 0}, false},
		{"800", false, []int{8, 0}, false},
		{"1200", true, []int{12, 0}, false},
		{"12:00", true, []int{12, 0}, false},
		{"1200", false, []int{12, 0}, false},
		{"12:00", false, []int{12, 0}, false},
		{"12", false, []int{12, 0}, false},
		{"12", false, []int{12, 0}, false},
		{"1200", true, []int{12, 0}, false},
		{"1200", false, []int{12, 0}, false},
		// Minutes
		{"0012", true, []int{0, 12}, false},
		{"00:12", true, []int{0, 12}, false},
		{"0012", false, []int{0, 12}, false},
		{"00:12", false, []int{0, 12}, false},
		{"008", true, []int{0, 8}, false},
		{"00:8", true, []int{0, 8}, false},
		{"008", false, []int{0, 8}, false},
		{"00:8", false, []int{0, 8}, false},
		{"0008", true, []int{0, 8}, false},
		{"00:08", true, []int{0, 8}, false},
		{"0008", false, []int{0, 8}, false},
		{"00:08", false, []int{0, 8}, false},
		{"012", true, []int{0, 12}, false},
		{"012", false, []int{0, 12}, false},
		{"12", true, []int{0, 12}, false},
		{"08", true, []int{0, 8}, false},
		{"8", true, []int{0, 8}, false},
		// Hours and minutes
		{"0812", true, []int{8, 12}, false},
		{"08:12", true, []int{8, 12}, false},
		{"0808", true, []int{8, 8}, false},
		{"08:08", true, []int{8, 8}, false},
		{"0812", false, []int{8, 12}, false},
		{"08:12", false, []int{8, 12}, false},
		{"0808", false, []int{8, 8}, false},
		{"08:08", false, []int{8, 8}, false},
		{"1212", true, []int{12, 12}, false},
		{"12:12", true, []int{12, 12}, false},
		{"1208", true, []int{12, 8}, false},
		{"12:08", true, []int{12, 8}, false},
		{"1212", false, []int{12, 12}, false},
		{"12:12", false, []int{12, 12}, false},
		{"1208", false, []int{12, 8}, false},
		{"12:08", false, []int{12, 8}, false},
		{"00:00", false, []int{0, 0}, false},
		{"0000", false, []int{0, 0}, false},
		// Wrong inputs
		{"12:08:00", false, nil, true},
		{"12:08:00", true, nil, true},
		{"12:1a", true, nil, true},
		{"1a:00", true, nil, true},
		{"12:1a", false, nil, true},
		{"1a:00", false, nil, true},
		{"121a", true, nil, true},
		{"1a00", true, nil, true},
		{"121a", false, nil, true},
		{"1a00", false, nil, true},
		{"39:00", true, nil, true},
		{"3900", true, nil, true},
		{"39:00", false, nil, true},
		{"3900", false, nil, true},
		{"00:72", true, nil, true},
		{"0072", true, nil, true},
		{"00:72", false, nil, true},
		{"0072", false, nil, true},
		{"072", false, nil, true},
		{"", false, nil, true},
		{"a", false, nil, true},
	}
	for _, tt := range testcases {
		name := "valid"
		if tt.wantErr {
			name = "invalid"
		}
		name += fmt.Sprintf(" with input %v, minutesFirst %v", tt.input, tt.minutesFirst)
		t.Run(name, func(t *testing.T) {
			got, err := HoursAndMinutes(tt.input, tt.minutesFirst)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"HoursAndMinutes(%q, %v) error = %q, wantErr %v",
					tt.input,
					tt.minutesFirst,
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"HoursAndMinutes(%q, %v) = %v, want %v",
					tt.input,
					tt.minutesFirst,
					got,
					tt.want,
				)
			}
		})
	}
}

func ExampleHoursAndMinutes() {
	hoursAndMinutes, _ := HoursAndMinutes("08:12", true)
	fmt.Println("hours:", hoursAndMinutes[0], "minutes:", hoursAndMinutes[1])
	// Output: hours: 8 minutes: 12
}
