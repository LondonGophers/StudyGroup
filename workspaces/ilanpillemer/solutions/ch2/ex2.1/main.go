package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	FreezingC Celsius = 0
	BoilingC  Celsius = 100
)

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CtoK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FtoK(f Fahrenheit) Kelvin  { return Kelvin(CtoK(FtoC(f))) }
func KtoC(k Kelvin) Celsius     { return Celsius(k - 273.15) }
func KtoF(k Kelvin) Fahrenheit  { return CtoF(KtoC(k)) }

func (c Celsius) String() string { return fmt.Sprintf("%.2fC", c) }
func (k Kelvin) String() string { return fmt.Sprintf("%.2fK", k) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2fF", f) }