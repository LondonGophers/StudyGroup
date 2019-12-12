// Package rotate is a version of `rotate` that operates in a single pass.
package rotate

// Left rotates a slice of ints 'n' indexes to the left in place with a single pass.
func Left(s []int, n int) {
	n %= len(s)
	count := 0

	for start := 0; count < len(s); start++ {
		current := start
		prev := s[start]

		for ok := true; ok; ok = start != current {
			next := (len(s) - n + current) % len(s)
			temp := s[next]
			s[next] = prev
			prev = temp
			current = next
			count++
		}
	}
}

// Right rotates a slice of ints 'n' indexes to the right in place with a single pass.
func Right(s []int, n int) {
	n %= len(s)
	count := 0

	for start := 0; count < len(s); start++ {
		current := start
		prev := s[start]

		for ok := true; ok; ok = start != current {
			next := (current + n) % len(s)
			temp := s[next]
			s[next] = prev
			prev = temp
			current = next
			count++
		}
	}
}
