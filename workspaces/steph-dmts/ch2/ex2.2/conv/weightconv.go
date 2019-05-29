package conv

import "fmt"
type Kilo float64
type Pound float64

const(
	KiloInP Kilo = 2.20462
	PoundInK Pound = 2.20462
)

func (k Kilo) String() string {
	return fmt.Sprintf("%g K",k)
}

func (p Pound) String() string{
	return fmt.Sprintf("%g P",p)
}