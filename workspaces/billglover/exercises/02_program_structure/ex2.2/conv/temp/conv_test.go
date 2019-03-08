package temp

import "testing"

var convCases = []struct {
	c Celsius
	f Fahrenheit
	k Kelvin
}{
	{c: -40, f: -40, k: 233.15},
	{c: 0, f: 32, k: 273.15},
	{c: 100, f: 212, k: 373.15},
}

func TestConversions(t *testing.T) {

	t.Run("from Celcius", func(st *testing.T) {
		for _, tc := range convCases {
			// Note: comparison casts to float32 to mask floating point innacuracy
			if got, want := float32(CtoF(tc.c)), float32(tc.f); got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}

			// Note: comparison casts to float32 to mask floating point innacuracy
			if got, want := float32(CtoK(tc.c)), float32(tc.k); got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("from Fahrenheit", func(st *testing.T) {
		for _, tc := range convCases {
			// Note: comparison casts to float32 to mask floating point innacuracy
			if got, want := float32(FtoC(tc.f)), float32(tc.c); got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}

			if got, want := float32(FtoK(tc.f)), float32(tc.k); got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("from Kelvin", func(st *testing.T) {
		for _, tc := range convCases {
			// Note: comparison casts to float32 to mask floating point innacuracy
			if got, want := float32(KtoC(tc.k)), float32(tc.c); got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}

			// Note: comparison casts to float32 to mask floating point innacuracy
			if got, want := float32(KtoF(tc.k)), float32(tc.f); got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

}

func TestConverter(t *testing.T) {
	cnv := Converter{}

	if got, want := cnv.Name(), "Temperature"; got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}

	txt, err := cnv.Convert(4)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got, want := txt, "4.00°C is 39.20°F, 277.15°K, 4.00°F is -15.56°C, 257.59°K, 4.00°K is -269.15°C, -452.47°F"; got != want {
		t.Errorf("\ngot:  %s\nwant: %s\n", got, want)
	}
}
