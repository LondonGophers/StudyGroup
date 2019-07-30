package weightconv

import "fmt"

type Pounds float64
type Kilograms float64

func (p Pounds) String() string    { return fmt.Sprintf("%g pound(s)", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%g Kg(s)", k) }
