// The instructor of the linear algebra course decides that calculus is now a
// prerequisite. Extend the topoSort function to report cycles.
package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	// create calculus -> linear algebra -> calculus cycle:
	"linear algebra": {"calculus"},

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

// topoSort traverses the directed graph, represented by m, visiting each node once.
// If a cycle is detected then a warning will be printed to stdout.
//
// topoSort uses a black, grey, and white list to show verticies visited (black),
// visiting (grey), and unvisited (white). When a vertex is encountered that is
// also being visited then we know we have a cycle in the graph.
func topoSort(m map[string][]string) []string {
	var order []string
	var visited = make(map[string]bool)  // aka `black`
	var visiting = make(map[string]bool) // aka `grey`

	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !visited[item] {
				if visiting[item] {
					fmt.Printf("WARNING: cycle detected from %s\n", item)
					return
				}
				visiting[item] = true

				adjacent := m[item]
				if len(adjacent) > 0 {
					visitAll(adjacent)
				} else {
					visiting[item] = false
				}

				visited[item] = true
				order = append(order, item)
			}
		}
	}

	var keys []string // aka `white`
	for key := range m {
		keys = append(keys, key)
	}

	visitAll(keys)
	return order
}
