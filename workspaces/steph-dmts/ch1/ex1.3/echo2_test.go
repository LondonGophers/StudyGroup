package bench

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func echo2() {
	fmt.Println(strings.Join(os.Args[1:], ""))
}
