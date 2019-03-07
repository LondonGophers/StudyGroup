package tempconv

// CtoF converts a temperature in Celsius to Fahrenheit
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CtoK converts a temperature in Celsius to Kelvin
func CtoK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// FtoC converts a temperature in Fahrenheit to Celsius
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FtoK converts a temperature in Fahrenheit to Kelvin
func FtoK(f Fahrenheit) Kelvin { return CtoK(FtoC(f)) }

// KtoC converts a temperature in Kelvin to Celsius
func KtoC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

// KtoF converts a temperature in Kelvin to Fahrenheit
func KtoF(k Kelvin) Fahrenheit { return CtoF(KtoC(k)) }
