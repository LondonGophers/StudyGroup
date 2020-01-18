// Package lengthconv performs conversion between length in feet and meters
package lengthconv

import "fmt"

type Feet float64
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%g ft", f) }
func (m Meters) String() string { return fmt.Sprintf("%g m", m) }
