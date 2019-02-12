// The type of each word used by `IntSet` is `uint64`, but 64-bit arithmetic
// may be inefficient on a 32-bit platform. Modify the program to use the `uint`
// type, which is the most efficient unsigned integer type for the platform.
// Instead of dividing by 64, define a constant holding the effective size of
// `uint` in bits, 32 or 64. You can use the perhaps too-clever expression
// `32 << (^uint(0) >> 63)` for this purpose.
package main

import (
	"fmt"
	"strings"
)

const size = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/size, uint(x%size)
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
	word, bit := x/size, uint(x%size)
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
	word, bit := x/size, uint(x%size)
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

// Elems returns a slice containing the elements of the set, suitable for
// iterating over with a range loop.
func (s *IntSet) Elems() []int {
	e := []int{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				e = append(e, size*i+j)
			}
		}
	}
	return e
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				if sb.Len() > len("{") {
					sb.WriteString(" ")
				}
				sb.WriteString(fmt.Sprintf("%d", size*i+j))
			}
		}
	}
	sb.WriteString("}")
	return sb.String()
}

func main() {
	var x IntSet
	x.AddAll(1, 2, 3, 5, 6, 7, 8, 9, 0)
	fmt.Printf("%s has 7: %t\n", x.String(), x.Has(7))
	fmt.Printf("%s has 4: %t\n", x.String(), x.Has(4))
}
