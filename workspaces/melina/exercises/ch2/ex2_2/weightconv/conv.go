package weightconv

// KToP converts Kgs to pounds.
func KToP(k Kilograms) Pounds { return Pounds(k *  2.205) }

// PToK converts Pounds to Kgs.
func PToK(f Pounds) Kilograms { return Kilograms(f / 2.205) }