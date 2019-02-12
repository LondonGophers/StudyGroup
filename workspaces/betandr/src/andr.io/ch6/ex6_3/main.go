// `(*IntSet).UnionWith` computes the union of two sets using `|`, the
// word-parallel bitwise OR operator. Implement methods for `IntersectWith`,
// `DifferenceWith`, and `SymmetricDifference` for the corresponding set
// operations. (The symmetric difference of two sets contains the elements
// present in one set or the other but not both.)
package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// AddAll adds the non-negative values x to the set.
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith computes the intersection between two sets
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWith computes the difference between two sets
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// SymmetricDifferenceWith computes the symmetric difference between two sets
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Len returns the number of elements by running a population count (by clearing)
// on each word.
func (s *IntSet) Len() int {
	count := 0
	for _, w := range s.words {
		for w != 0 {
			w = w & (w - 1) // clear rightmost non-zero bit
			count++
		}
	}
	return count
}

// Remove removes `x` from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] ^= 1 << bit
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = nil // is this a cheat? :thinking_face:
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	ns := new(IntSet)
	for _, word := range s.words {
		ns.words = append(ns.words, word)
	}
	return ns
}

// String returns the set as a string of the form "{1 2 3}".
// TODO re-write using strings.Builder?
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.AddAll(1, 2)
	y.AddAll(2, 3)
	fmt.Printf("set %s has %d items\n", x.String(), x.Len())
	fmt.Printf("set %s has %d items\n", y.String(), y.Len())

	// set union
	x.UnionWith(&y)
	fmt.Printf("%s union has %d items\n", x.String(), x.Len())

	x.Clear()
	y.Clear()
	x.AddAll(1, 2)
	y.AddAll(2, 3)

	// set intersection
	x.IntersectWith(&y)
	fmt.Printf("%s intersection has %d items\n", x.String(), x.Len())

	x.Clear()
	y.Clear()
	x.AddAll(1, 2)
	y.AddAll(2, 3)

	// set difference
	x.DifferenceWith(&y)
	fmt.Printf("%s difference has %d items\n", x.String(), x.Len())

	x.Clear()
	y.Clear()
	x.AddAll(1, 2)
	y.AddAll(2, 3)

	// set symmetric difference
	x.SymmetricDifferenceWith(&y)
	fmt.Printf("%s symmetric difference has %d items\n", x.String(), x.Len())
}
