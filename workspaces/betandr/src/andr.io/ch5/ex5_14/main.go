// Use the `breadthFirst` function to explore a different structure. For example,
// you could use the course dependencies from the `topoSort` example (a
// directed graph), the file system hierarchy on your computer (a tree), or a
// list of bus or subway routes downloaded from your city government’s web site
// (an undirected graph).
package main

import (
	"fmt"
)

// directedGraph is the prereqs map from §5.11
var directedGraph = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// crawl returns the child verticies from `directedGraph` for a given key.
func crawl(item string) []string {
	fmt.Println(item)
	list := directedGraph[item]
	return list
}

func main() {
	var keys []string
	for key := range directedGraph {
		keys = append(keys, key)
	}

	// depth-first traverse of the keys from directedGraph
	breadthFirst(crawl, keys)
}
