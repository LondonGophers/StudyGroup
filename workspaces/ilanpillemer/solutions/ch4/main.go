package main

//DIffCount simply compares every bit of 256
func DiffCount(left [32]byte, right [32]byte) int {
	n := 0
	for i := 0; i < 32; i++ {
		for j := uint(0); j < 8; j++ {
			l, r := left[i]&byte(1<<j), right[i]&byte(1<<j)
			if l != r {
				n++
			}
		}
	}
	return n
}

//$ go test
//The differing bits of the SHA256 of x and X is 125.
//PASS
//ok  	_/Users/ilanpillemer/git/study-group/workspaces/ilanpillemer/solutions/ch4	0.005s
//$

