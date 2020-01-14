package puzzleutil

import "testing"

func TestCalculateFuel(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{12}, 2},
		{[]int{14}, 2},
		{[]int{1969}, 966},
		{[]int{100756}, 50346},
	}

	for _, c := range cases {
		got := CalculateFuel(c.in)
		if got != c.want {
			t.Errorf("CalculateFuel(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
