package main

import(
	"os"
	"fmt"
)
//Prints command and its args in brackets, as an array
// func main()  {
// 	fmt.Println(os.Args[:])
// }

//Prints command and its args separated by a space 
func main()  {
	for _, arg := range os.Args[:] {
		fmt.Print(arg+" ")
	}
}