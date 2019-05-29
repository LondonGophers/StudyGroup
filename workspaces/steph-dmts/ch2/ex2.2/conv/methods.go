package conv


func KToP(k Kilo)Pound  {
	return Pound(k*KiloInP)
}
func PToK(p Pound)Kilo{
	return Kilo(p*PoundInK)
}

func MToF(m Meter)Foot  {
	return Foot(m*FinM)
}
func FToM(f Foot)Meter{
	return Meter(f/MinF)
}