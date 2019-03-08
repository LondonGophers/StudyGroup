// Package temp performs temperature conversions.
package temp

import "fmt"

// Celsius is the temperature in degrees Celsius.
type Celsius float64

// Fahrenheit is the temperature in degrees Fahrenheit.
type Fahrenheit float64

// Kelvin is the temperature in degrees Kelvin.
type Kelvin float64

const (
	// AbsoluteZeroC is absolute zero in °C.
	AbsoluteZeroC Celsius = -273.15

	// FreezingC is the freezing point of water in °C.
	FreezingC Celsius = 0

	// BoilingC is the boiling point of water in °C.
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%.2f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.2f°K", k) }
