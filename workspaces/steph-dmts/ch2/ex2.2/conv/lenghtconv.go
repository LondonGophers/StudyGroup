package conv

import "fmt"

type Meter float64
type Foot float64

const(
	MinF = 3.28084
	FinM = 1/3.28084
)

func (m Meter) String() string{
	return fmt.Sprintf("%g M",m)
}

func (f Foot) String() string{
	return fmt.Sprintf("%g F",f)
}