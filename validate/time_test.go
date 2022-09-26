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
