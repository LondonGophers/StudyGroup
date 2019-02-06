package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Printf("index[%d] value:[%s]\n", i, v)
	}
}

//go run main.go 1 2 3
//index[0] value:[/var/folders/cw/6_jr6ms12wgchzxj9s9kq5sm0000gn/T/go-build652312564/b001/exe/main]
//index[1] value:[1]
//index[2] value:[2]
//index[3] value:[3]
