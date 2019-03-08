package length

import "fmt"

// Converter implements the Convertable interface
type Converter struct{}

// Name returns the common name for the collaction of units in this package
func (cnv *Converter) Name() string {
	return "Length"
}

// Convert returns a string representation of the provided float in various units
func (cnv *Converter) Convert(v float64) (string, error) {
	i := Inches(v)
	txt := fmt.Sprintf("%s is %s", i, ItoM(i))

	mm := Millimeters(v)
	txt += fmt.Sprintf(", %s is %s", mm, MtoI(mm))

	return txt, nil
}

// ItoM converts a length in Inches to Millimeteres
func ItoM(i Inches) Millimeters { return Millimeters(i * MilimetersPerInch) }

// MtoI converts a length in Millimeteres to Inches
func MtoI(mm Millimeters) Inches { return Inches(mm / MilimetersPerInch) }
