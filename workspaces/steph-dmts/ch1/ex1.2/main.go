package main

import(
	"os"
	"fmt"
)

//Prints index and value of each arg on a new line
func main()  {
	for index,arg := range os.Args[1:] {
		fmt.Println(index,arg)
	}
}