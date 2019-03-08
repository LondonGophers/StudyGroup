package distance

import "fmt"

// Converter implements the Convertable interface
type Converter struct{}

// Name returns the common name for the collaction of units in this package
func (cnv *Converter) Name() string {
	return "Distance"
}

// Convert returns a string representation of the provided float in various units
func (cnv *Converter) Convert(v float64) (string, error) {
	m := Miles(v)
	txt := fmt.Sprintf("%s is %s", m, MtoK(m))

	km := Kilometers(v)
	txt += fmt.Sprintf(", %s is %s", km, KtoM(km))

	return txt, nil
}

// MtoK converts a length in Miles to Kilometers
func MtoK(m Miles) Kilometers { return Kilometers(m / MilesPerKilometer) }

// KtoM converts a length in Kilometers to Miles
func KtoM(km Kilometers) Miles { return Miles(km * MilesPerKilometer) }
