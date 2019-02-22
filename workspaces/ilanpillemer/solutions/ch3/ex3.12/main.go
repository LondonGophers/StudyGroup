package main

import (
	"sort"
	"strings"
)

func anagram(s1 string, s2 string) bool {
	a1, a2 := strings.Split(s1, ""), strings.Split(s2, "")
	sort.Strings(a1)
	sort.Strings(a2)
	return strings.Join(a1, "") == strings.Join(a2, "")
}
