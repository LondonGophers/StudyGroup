package adjacent_test

import (
	"fmt"
	"sort"

	adjacent "github.com/go-london-user-group/study-group/workspaces/jlucktay/ch4/ex4_5"
)

func ExampleEliminate() {
	s := []string{"one", "one", "two", "two", "two", "three", "three", "three", "three"}
	adjacent.Eliminate(&s)
	fmt.Printf("%#v\n", s)
	// Output: []string{"one", "two", "three"}
}

func ExampleEliminate_empty_slice() {
	s := []string{}
	adjacent.Eliminate(&s)
	fmt.Printf("%#v\n", s)
	// Output: []string{}
}

func ExampleEliminate_duplicate_words() {
	s := []string{"Oh", "we", "we", "all", "we", "all", "all", "all", "become", "one", "one", "one", "one"}
	adjacent.Eliminate(&s)
	fmt.Printf("%#v\n", s)
	// Output: []string{"Oh", "we", "all", "we", "all", "become", "one"}
}

func ExampleEliminate_numbers() {
	s := []string{"3", "2", "1", "4", "3", "2", "1", "4", "1"}
	sort.Strings(s)
	adjacent.Eliminate(&s)
	fmt.Printf("%#v\n", s)
	// Output: []string{"1", "2", "3", "4"}
}
