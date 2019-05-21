package bench

import(
	"os"
	"fmt"
	"testing"
	
)


func BenchmarkEcho(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func echo1(){
	var s,sep string
	for i := 0; i < len(os.Args); i++ {
		s+=sep+os.Args[i]
		sep=" "
	}
	fmt.Println(s)
}

