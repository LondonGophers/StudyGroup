// Write a `String` method for the `*tree` type in `gopl.io/ch4/treesort` (§4.4) that
// reveals the sequence of values in the tree.
package treesort

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

// String returns an pre-order representation of the contents of a tree
func (s *tree) String() string {
	sb := new(strings.Builder)
	sb.WriteString("[")
	var traverse func(t *tree)
	traverse = func(t *tree) {
		if t == nil {
			return
		}
		sb.WriteString(fmt.Sprintf("%d ", t.value))
		traverse(t.left)
		traverse(t.right)
	}
	traverse(s)
	sb.WriteString("]")
	return sb.String()
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// Render returns a string representation of a tree
func Render(values []int) string {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root.String()
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
