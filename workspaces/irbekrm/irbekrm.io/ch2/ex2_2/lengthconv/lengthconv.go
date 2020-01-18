package lengthconv

// FToM converts length in feet to length in meters
func FtToM(f Feet) Meters { return Meters(f / 3.2808) }

// MToF converts length in meters to length in feet
func MToFt(m Meters) Feet { return Feet(m * 3.2808) }
