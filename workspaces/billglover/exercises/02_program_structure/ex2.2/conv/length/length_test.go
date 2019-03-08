package length

import "testing"

var stringerCases = []struct {
	l float64
	i string
	m string
}{
	{l: 0, i: "0.00 inches", m: "0.00 mm"},
	{l: 100, i: "100.00 inches", m: "100.00 mm"},
	{l: -10.5, i: "-10.50 inches", m: "-10.50 mm"},
}

func TestString(t *testing.T) {

	t.Run("Inches to String", func(st *testing.T) {
		for _, tc := range stringerCases {
			if got, want := Inches(tc.l).String(), tc.i; got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("Milimeters to String", func(st *testing.T) {
		for _, tc := range stringerCases {
			if got, want := Millimeters(tc.l).String(), tc.m; got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})
}
