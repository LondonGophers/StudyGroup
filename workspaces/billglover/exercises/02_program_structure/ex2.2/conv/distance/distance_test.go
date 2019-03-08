package distance

import "testing"

var stringerCases = []struct {
	d  float64
	km string
	mi string
}{
	{d: 0, km: "0.00 km", mi: "0.00 mi"},
	{d: 100, km: "100.00 km", mi: "100.00 mi"},
	{d: -10.5, km: "-10.50 km", mi: "-10.50 mi"},
}

func TestString(t *testing.T) {

	t.Run("Kilometers to String", func(st *testing.T) {
		for _, tc := range stringerCases {
			if got, want := Kilometers(tc.d).String(), tc.km; got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("Miles to String", func(st *testing.T) {
		for _, tc := range stringerCases {
			if got, want := Miles(tc.d).String(), tc.mi; got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})
}
