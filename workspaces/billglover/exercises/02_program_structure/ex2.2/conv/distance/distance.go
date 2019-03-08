// Package distance performs distance conversions
package distance

import "fmt"

const (
	// MilesPerKilometer is the number of miles per kilometer
	MilesPerKilometer = 0.62137
)

// Miles is the temperature in degrees Miles.
type Miles float64

// Kilometers is the temperature in degrees Kilometers.
type Kilometers float64

func (m Miles) String() string       { return fmt.Sprintf("%.2f mi", m) }
func (km Kilometers) String() string { return fmt.Sprintf("%.2f km", km) }
