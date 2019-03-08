package length

import "testing"

var convCases = []struct {
	i Inches
	m Millimeters
}{
	{i: -40, m: -1016},
	{i: 0, m: 0},
	{i: 100, m: 2540},
}

func TestConversions(t *testing.T) {

	t.Run("from Inches", func(st *testing.T) {
		for _, tc := range convCases {
			// Note: comparison casts to float32 to mask floating point innacuracy
			if got, want := float32(ItoM(tc.i)), float32(tc.m); got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})

	t.Run("from Milimeters", func(st *testing.T) {
		for _, tc := range convCases {
			// Note: comparison casts to float32 to mask floating point innacuracy
			if got, want := float32(MtoI(tc.m)), float32(tc.i); got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	})
}

func TestConverter(t *testing.T) {
	cnv := Converter{}

	if got, want := cnv.Name(), "Length"; got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}

	txt, err := cnv.Convert(4)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if got, want := txt, "4.00 inches is 101.60 mm, 4.00 mm is 0.16 inches"; got != want {
		t.Errorf("\ngot:  %s\nwant: %s\n", got, want)
	}
}
