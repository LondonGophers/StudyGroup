// Write a general purpose unit-conversion program analogous to `cf` that reads
// numbers from its command-line arguments or from the standard imput if there are
// no arguments, and converts each nunber into units like temperature in Celsius
// and Fahrenheit, length in feet and meters, weight in pounds, kilograms, and the
// like.
package weightconv

import "fmt"

// Kilo weight type
type Kilo float64

// Stone weight type
type Stone float64

// Pound weight type
type Pound float64

// Conversion constants
const (
	KiloP Kilo = 2.2046226218
	KiloS Kilo = 6.35029318

	PoundK Pound = 0.453592
	PoundS Pound = 0.0714286

	StoneK Stone = 0.15747
	StoneP Stone = 14
)

// Adds custom String() formatting to Kilo
func (f Kilo) String() string { return fmt.Sprintf("%gkg", f) }

// Adds custom String() formatting to Stone
func (f Stone) String() string { return fmt.Sprintf("%gst", f) }

// Adds custom String() formatting to Pound
func (f Pound) String() string { return fmt.Sprintf("%glb", f) }
