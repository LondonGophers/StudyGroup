package distconv

import "fmt"

type Mile float64
type Kilometer float64
type NauticalMile float64
type RomanMile float64
type ChineseMile float64

func MiToKi(mi Mile) Kilometer              { return Kilometer(mi * 1.609344) }
func KiToMi(km Kilometer) Mile              { return Mile(km / 1.609344) }
func NautMiToKi(nmi NauticalMile) Kilometer { return Kilometer(nmi * 1.852) }
func KiToNautMi(km Kilometer) NauticalMile  { return NauticalMile(km / 1.852) }
func ChiMiToKi(cmi ChineseMile) Kilometer   { return Kilometer(cmi * 0.5) }
func KiToChiMi(km Kilometer) ChineseMile    { return ChineseMile(km / 0.5) }

func MiToNautMi(mi Mile) NauticalMile { return KiToNautMi(MiToKi(mi)) }
func MiToChiMi(mi Mile) ChineseMile  { return KiToChiMi(MiToKi(mi)) }

func (km Kilometer) String() string     { return fmt.Sprintf("%.2fkm", km) }
func (mi Mile) String() string          { return fmt.Sprintf("%.2fmi", mi) }
func (nmi NauticalMile) String() string { return fmt.Sprintf("%.2f Nautical mi", nmi) }
func (cmi ChineseMile) String() string  { return fmt.Sprintf("%.2f Chinese mi", cmi) }
