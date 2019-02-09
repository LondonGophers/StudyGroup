// Write a general purpose unit-conversion program analogous to `cf` that reads
// numbers from its command-line arguments or from the standard imput if there are
// no arguments, and converts each nunber into units like temperature in Celsius
// and Fahrenheit, length in feet and meters, weight in pounds, kilograms, and the
// like.
package weightconv

// KToP converts a Kilo weight to Pounds.
func KToP(k Kilo) Pound { return Pound(k * KiloP) }

// KToS converts a Kilo temperature to Stone.
func KToS(k Kilo) Stone { return Stone(k / KiloS) }

// PToK converts a Pound weight to Kilo.
func PToK(p Pound) Kilo { return Kilo(p * PoundK) }

// PToS converts a Pound temperature to Stone.
func PToS(p Pound) Stone { return Stone(p * PoundS) }

// SToK converts a Stone weight to Kilo.
func SToK(s Stone) Kilo { return Kilo(s / StoneK) }

// SToP converts a Stone temperature to Pound.
func SToP(s Stone) Pound { return Pound(s * StoneP) }
