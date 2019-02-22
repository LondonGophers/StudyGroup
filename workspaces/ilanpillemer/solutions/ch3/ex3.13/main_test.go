package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("KB", KB)
	fmt.Println("MB", MB)
	fmt.Println("GB", GB)
	fmt.Println("TB", TB)
	fmt.Println("PB", PB)
	fmt.Println("EB", EB)
	fmt.Println("ZB", ZB)
	fmt.Println("YB", YB)
	fmt.Println()
}

func Test1024(t *testing.T) {
	fmt.Println("KiB", KiB)
	fmt.Println("MiB", MiB)
	fmt.Println("GiB", GiB)
	fmt.Println("TiB", TiB)
	fmt.Println("PiB", PiB)
	fmt.Println("EiB", EiB)
	fmt.Println("ZiB", float64(ZiB))
	fmt.Println("YiB", float64(YiB))

}