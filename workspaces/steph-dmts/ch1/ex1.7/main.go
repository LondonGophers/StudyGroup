package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
//io.Copy does not give an error
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr,"fetch: %v\n",err)
			os.Exit(1)
		}
		io.Copy(os.Stdout,resp.Body)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
		// }
		// resp.Body.Close()
		// fmt.Printf("%v",body)
	}
}
