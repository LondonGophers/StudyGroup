// Write a `String` method for the `*tree` type in `gopl.io/ch4/treesort` (§4.4) that
// reveals the sequence of values in the tree.
package main

import (
	"fmt"
	"math/rand"
	"sort"

	"andr.io/ch7/ex7_3/treesort"
)

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	fmt.Println(treesort.Render(data))
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		fmt.Printf("not sorted: %v", data)
	}
}
