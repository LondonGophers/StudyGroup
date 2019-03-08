// Package length performs length conversions
package length

import "fmt"

const (
	// MilimetersPerInch is the number of milimeters per inch
	MilimetersPerInch = 25.4
)

// Inches is the temperature in degrees Inches.
type Inches float64

// Millimeters is the temperature in degrees Millimeters.
type Millimeters float64

func (i Inches) String() string       { return fmt.Sprintf("%.2f inches", i) }
func (mm Millimeters) String() string { return fmt.Sprintf("%.2f mm", mm) }
