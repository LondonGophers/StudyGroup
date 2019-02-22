package main

// you cannot use functions as the constants are untyped, you can only use operators
// there is no exponent operator

const (
	KB = 1E3
	MB = 1E6
	GB = 1E9
	TB = 1E12
	PB = 1E15
	EB = 1E18
	ZB = 1E21
	YB = 1E24
)

const (
	_ = 1 << (iota * 10)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

//$ go test
//KB 1000
//MB 1000000
//GB 1000000000
//TB 1000000000000
//PB 1000000000000000
//EB 1000000000000000000
//ZB 1e+21
//YB 1e+24
//
//KiB 1024
//MiB 1048576
//GiB 1073741824
//TiB 1099511627776
//PiB 1125899906842624
//EiB 1152921504606846976
//ZiB 1.1805916207174113e+21
//YiB 1.2089258196146292e+24
//PASS
//ok  	_/Users/ilanpillemer/git/study-group/workspaces/ilanpillemer/solutions/ch3/ex3.13	0.005s
//$
