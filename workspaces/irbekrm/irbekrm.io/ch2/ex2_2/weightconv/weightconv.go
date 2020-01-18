// Package performs conversions between weight in pounds and in kilograms
package weightconv

import "fmt"

type Pounds float64
type Kilograms float64

func (p Pounds) String() string    { return fmt.Sprintf("%g lbs", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%g kg", k) }
