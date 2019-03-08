package distance

import "testing"

var convCases = []struct {
	mi Miles
	km Kilometers
}{
	{mi: -40, km: -64.373886},
	{mi: 0, km: 0},
	{mi: 100, km: 160.93471},
}

func TestConversions(t *testing.T) {

	t.Run("from Kilometers", func(st *testing.T) {
		for _, tc := range convCases {
			// Note: comparison casts to float32 to mask floating point innacuracy
			if got, want := float32(KtoM(tc.km)), float32(tc.mi); got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("from Miles", func(st *testing.T) {
		for _, tc := range convCases {
			// Note: comparison casts to float32 to mask floating point innacuracy
			if got, want := float32(MtoK(tc.mi)), float32(tc.km); got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})
}

func TestConverter(t *testing.T) {
	cnv := Converter{}

	if got, want := cnv.Name(), "Distance"; got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}

	txt, err := cnv.Convert(4)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got, want := txt, "4.00 mi is 101.60 km, 4.00 km is 0.16 mi"; got != want {
		t.Errorf("\ngot:  %s\nwant: %s\n", got, want)
	}
}
