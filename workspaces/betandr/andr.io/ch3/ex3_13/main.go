// Write const declarations for KB, MB, up through YB as compactly as you can.
package main

import "fmt"

const (
	_   = 1 << (10 * iota)
	KiB // 2^10
	MiB // 2^20
	GiB // 2^30
	TiB // 2^40
	PiB // 2^50
	EiB // 2^60
	ZiB // 2^70
	YiB // 2^80
)

const (
	KB = 1000      // 10^3
	MB = KB * 1000 // 10^6
	GB = MB * 1000 // 10^9
	TB = GB * 1000 // 10^12
	PB = TB * 1000 // 10^15
	EB = PB * 1000 // 10^18
	ZB = EB * 1000 // 10^21
	YB = ZB * 1000 // 10^24
)

func main() {
	fmt.Printf("%d\n", EiB)
	fmt.Printf("%d\n", EB)
}
