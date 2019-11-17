package tempconv

// CELSIUS TO ...
// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Clesius) Kelvin { return Kelvin(c + KelvinC) }

// FAHRENHEIT TO ...
// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin(((f - 32) * 5 / 9) + KelvinC) }

// KELVIN TO ...
// KToC converts a Kelvin temperature to Celcius
func KToC(k Kelvin) Celsius { return Celsius(k - KelvinC) }

// KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-KelvinC)*9/5 + 32) }
