package lengthconv

import "fmt"

type Feet float64
type Meters float64

// const (
// 	AbsoluteZeroC Celsius = -273.15
// 	FreezingC     Celsius = 0
// 	BoilingC      Celsius = 100
// )

func (f Feet) String() string   { return fmt.Sprintf("%g feet", f) }
func (m Meters) String() string { return fmt.Sprintf("%g meter(s)", m) }
