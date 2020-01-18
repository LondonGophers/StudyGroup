package weightconv

// PtoKg converts weight in pounds to weight in kilograms
func PToKg(p Pounds) Kilograms { return Kilograms(p * 0.45359237) }

// KgtoP converts weight in kilograms to weight in pounds
func KgToP(k Kilograms) Pounds { return Pounds(k / 0.45359237) }
