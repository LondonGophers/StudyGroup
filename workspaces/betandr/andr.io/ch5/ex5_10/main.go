// Rewrite `toposort` to use maps instead of slices and eliminate the initial
// sort. Verify that the results, though nondeterministic, are valid
// topological orderings.
package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
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

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// topoSort traverses the directed graph represented by `m` visiting each node once.
func topoSort(m map[string][]string) []string {
	var order []string
	var visited = make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !visited[item] {
				visited[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	visitAll(keys)
	return order
}
