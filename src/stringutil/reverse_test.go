package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, expectedResult string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"H", "H"},
		{"", ""},
	}
	for _, c := range cases {
		actualResult := Reverse(c.in)
		if actualResult != c.expectedResult {
			t.Errorf("Reverse(%q) != %q, expected %q", c.in, actualResult, c.expectedResult)
		}
	}
}
