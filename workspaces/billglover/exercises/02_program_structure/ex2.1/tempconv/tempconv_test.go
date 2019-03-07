package tempconv

import "testing"

var stringerCases = []struct {
	t float64
	f string
	c string
	k string
}{
	{t: 0, c: "0°C", f: "0°F", k: "0°K"},
	{t: 100, c: "100°C", f: "100°F", k: "100°K"},
	{t: -10.5, c: "-10.5°C", f: "-10.5°F", k: "-10.5°K"},
	{t: -1468.141592657, c: "-1468.141592657°C", f: "-1468.141592657°F", k: "-1468.141592657°K"},
}

func TestString(t *testing.T) {

	t.Run("Celsius to String", func(st *testing.T) {
		for _, tc := range stringerCases {
			if got, want := Celsius(tc.t).String(), tc.c; got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("Fahrenheit to String", func(st *testing.T) {
		for _, tc := range stringerCases {
			if got, want := Fahrenheit(tc.t).String(), tc.f; got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("Kelvin to String", func(st *testing.T) {
		for _, tc := range stringerCases {
			if got, want := Kelvin(tc.t).String(), tc.k; got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})
}
