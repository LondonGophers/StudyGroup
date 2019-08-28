// ### Exercise 1.1
// Modify the `echo` program to also allow print `os.Args[0]`, the name of the
// command that invoked it.
package main
import (
	"fmt"
	"os"
)

func main(){
	var s, sep string 
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
	for i := 1; i<len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " " 
	}
	fmt.Println(s)
}