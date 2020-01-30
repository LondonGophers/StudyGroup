# Chapter 13 - Low-Level Programming

## Example: Deep Equivalence

### Exercise 13.1
Define a deep comparison function that considers numbers (of any type) equal if
they differ by less than one part in a billion.

### Exercise 13.3
Write a function that reports whether its argument is a cyclic data structure.

### Exercise 13.4
Use `sync.Mutex` to make `bzip2.Writer` safe for concurrent use by multiple
goroutines.

### Exercise 13.5
Depending on C libraries has its drawbacks. Provide an alternative pure-Go
implementation of `bzip.NewWriter` that uses the `os/exec` package to run
`/bin/bzip2` as a subprocess.
