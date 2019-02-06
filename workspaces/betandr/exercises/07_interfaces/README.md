# Chapter 7 - Interfaces

## Interfaces as Contracts

### Exercise 7.1
Using the ideas from `ByteCounter`, implement counters for words and for lines.
You will find `bufio.ScanWords` useful.

### Exercise 7.2
Write a function `CountingWriter` with the signature below that, given an
`io.Writer`, returns a new `Writer` that wraps the original, and a pointer to an
`int64` variable that any moment contains the number of bytes written to the new
`Writer`.
```
  func CounterWriter(w io.Writer) (io.Writer, *int64)
```

### Exercise 7.3
Write a `String` method for the `*tree` type in `gopl.io/ch4/treesort` (§4.4) that
reveals the sequence of values in the tree.

## Interface Types

### Exercise 7.4
The `strings.NewReader` function returns a value that satisfies the `io.Reader`
interface (and others) by reading from its argument, a string. Implement a
simple version of `NewReader` yourself, and use it to make the HTML parser (§5.2)
take input from a string.

### Exercise 7.5
The `LimitReader` function in the `io` package accepts an `io.Reader r` and a
number of bytes `n`, and returns another `Reader` that reads from `r` but
reports an end-of-file condition after `n` bytes. Implement it.
```
  func LimitReader(r, io.Reader, n int64) io.Reader
```

## Parsing Flags with `flag.Value`

### Exercise 7.6
Add support for Kelvin temperatures to `tempFlag`.

### Exercise 7.7
Explain why the help message contains °C when the default value of
`20.0` does not.

## Sorting with `sort.Interface`

### Exercise 7.8
Many GUIs provide a table widget with a stateful multi-tier sort: the primary
sort key is the most recently clicked column head, the secondary sort key is the
second most clicked column head, and so on. Define an implementation of
`sort.Interface` for use by such a table. Compare that approach with repeated
sorting using `sort.Stable`.

### Exercise 7.9
Use the `html/template` package (§4.6) to replace `printTracks` with a function
that displays the tracks as an HTML table. Use the solution to the previous
exercise to arrange that each click on a column head makes an HTTP request to
sort the table.

### Exercise 7.10
The `sort.Interface` type can be adapted to other uses. Write a function
`IsPalindrome(s sort.Interface) bool` that reports whether the sequence `s` is a
palindrome, in other words, reversing the sequence would not change it. Assume
that the elements at indices `i` and `j` are equal if
`!s.Less(i, j) && !s.Less(j, i)`.

## The `http.Handler` Interface

### Exercise 7.11
Add additional handlers so that clients can create, read, update, and delete
database entries. For example, a request of the form
`/update?item=socks&price=6` will update the price of an item in the inventory
and report an error if the item does not exist or if the price is invalid.
(Warning: this change introduces concurrent variable updates.)

### Exercise 7.12
Change the handler for `/list` to print its output as an HTML table, not text.
You may find the `html/template` package (§4.6) useful.

## Example: Expression Evaluator

### Exercise 7.13
Add a `String` method to `Expr` to pretty-print the syntax tree. Check that the
results, when parsed again, yield an equivalent tree.

### Exercise 7.14
Define a new concrete type that satisfies the `Expr` interface and provides a
new operation such as computing the minimum value of its operands. Since the
`Parse` function does not create instances of this new type, to use it you will
need to construct a syntax tree directly (or extend the parser).

### Exercise 7.15
Write a program that reads a single expression from the standard input, prompts
the user to provide values for any variables, then evaluates the expression in
the resulting environment. Handle all errors gracefully.

### Exercise 7.16
Write a web-based calculator program.

## Example: Token-Based XML Decoding

### Exercise 7.17
Extend `xmlselect` so that elements may be selected not just by name, but by
their attributes too, in the manner of CSS, so that, for instance, an element
like `<div id="page" class="wide">` could be selected by matching `id` or
`class` as well as its name.

### Exercise 7.18
Using the token-based decoder API, write a program that will read an arbitrary
XML document and construct a tree of generic nodes that represents it. Nodes are
of two kinds: `CharData` nodes represent text strings, and `Element` nodes
represent named elements and their attributes. Each element node has a slice of
child nodes.

You may find the following declarations helpful.
```
  import "encoding/xml"

  type Node interface{} // CharData or *Element

  type CharData string

  type Element struct {
    Type      xml.Name
    Attr      []xml.Attr
    Children  []Node
  }
```
