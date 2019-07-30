package lengthconv

// MToF converts Meters to feet.
func MToF(m Meters) Feet { return Feet(m * 3.281) }

// FToM converts Feet to Meters.
func FToM(f Feet) Meters { return Meters(f / 3.281) }
