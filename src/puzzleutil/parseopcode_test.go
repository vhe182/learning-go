package puzzleutil

import (
	"reflect"
	"testing"
)

func TestParseOpCode(t *testing.T) {
	cases := []struct {
		in   []string
		want []int
	}{
		{[]string{"1", "0", "0", "0", "99"}, []int{2, 0, 0, 0, 99}},
		{[]string{"2", "3", "0", "3", "99"}, []int{2, 3, 0, 6, 99}},
		{[]string{"2", "4", "4", "5", "99", "0"}, []int{2, 4, 4, 5, 99, 9801}},
		{[]string{"1", "1", "1", "4", "99", "5", "6", "0", "99"}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, c := range cases {
		got := ParseOpCode(c.in)

		// can't do this comparison.  Figure out how to compare entire slices
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("ParseOpCode(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}
