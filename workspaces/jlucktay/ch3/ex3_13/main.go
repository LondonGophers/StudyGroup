// Write `const` declarations for KB, MB, up through YB as compactly as you can.
package main

import "fmt"

const (
	kB = 1e3
	MB = 1e6
	GB = 1e9
	TB = 1e12
	PB = 1e15
	EB = 1e18
	ZB = 1e21
	YB = 1e24
)

func main() {
	fmt.Printf("kB: %v\n", kB)
	fmt.Printf("MB: %v\n", MB)
	fmt.Printf("GB: %v\n", GB)
	fmt.Printf("TB: %v\n", TB)
	fmt.Printf("PB: %v\n", PB)
	fmt.Printf("EB: %v\n", EB)
	fmt.Printf("ZB: %v\n", ZB)
	fmt.Printf("YB: %v\n", YB)
}
