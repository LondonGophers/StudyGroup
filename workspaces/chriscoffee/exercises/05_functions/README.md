# Chapter 5 - Functions

## Recursion

### Exercise 5.1
Change the `findLinks` program to traverse the `n.FirstChild` linked list using
recursive calls to `visit` instead of a loop.

### Exercise 5.2
Write a function to populate a mapping from element names—`p`,
`div`, `span`, and so on—to the number of elements with that name in an HTML
document tree.

### Exercise 5.3
Write a function to print the contents of all text nodes in an HTML document tree.
Do not descend into `<script>` or `<style>` elements, since their contents are
not visible in a web browser.

### Exercise 5.4
Extend the `visit` function so that it extracts other kinds of links from the
document, such as images, scripts, and style sheets.

## Multiple Return Values

### Exercise 5.5
Implement `countWordsAndImages`. See (Exercise 4.9 for word-splitting).

### Exercise 5.6
Modify the `corner` function in `gopl.io/ch3/surface` (3.2) to use named results
and a bare return statement.

## Function Values

### Exercise 5.7
Develop `startElement` and `endElement` into a general HTML pretty-printer.
Print comment nodes, text nodes, and the attributes of each element
(<a href='...'). Use short forms like `<img/>` instead of `<img></img>` when an
element has no children. Write a test to ensure that the output can be parsed
successfully. (See Chapter 11.)

### Exercise 5.8
Modify `forEachNode` so that the `pre` and `post` functions return a boolean
result indicating whether to continue the traversal. Use it to write a function
`ElementByID` with the following signature that finds the first HTML element
with the specified `id` attribute. The function should stop the traversal as
soon as a match is found.
```
  func ElementById(doc *html.Node, id string) *html.Node
```

### Exercise 5.9
Write a function `expand(s string, f func(string) string) string` that replaces
each substring `"$foo"` within `s` by the text returned by `f("foo")`.

## Anonymous Functions

### Exercise 5.10
Rewrite `topoSort` to use maps instead of slices and eliminate the initial sort.
Verify that the results, through nondeterministic, are valid topological
orderings.

### Exercise 5.11
The instructor of the linear algebra course decides that calculus is now a
prerequisite. Extend the `topoSort` function to report cycles.

### Exercise 5.12
The `startElement` and `endElement` functions in `gopl.io/ch5/outline2` (§5.5)
share a global variable, `depth`. Turn them into anonymous functions that share
a variable local to the `outline` function.

### Exercise 5.13
Modify `crawl` to make local copies of the pages it finds, creating directories
as necessary. Don't make copies of pages that come from a different domain. For
example, if the original page comes from `golang.org`, save all files from
there, but exclude ones from `vimeo.com`.

### Exercise 5.14
Use the `breadFirst` function to explore a different structure. For example, you
could use the course dependencies from the `topoSort` example (a directed graph),
the file system hierarchy on your computer (a tree), or a list of bus or subway
routes downloaded from your city government's website (an undirected graph).

### Exercise 5.15
Write variadic functions `max` and `min`, analogous to `sum`. What should these
functions do when called with no arguments? Write variants that require at least
one argument.

### Exercise 5.16
Write variadic versions of `strings.Join`.

### Exercise 5.17
Write a variadic function `ElementsByTagName` that, given an HTML node tree and
zero or more names, returns all elements that match one of those names. Here are
two example calls:
```
  func ElementsByTagName(doc *html.Node, name ...string) []*html.Node

  images := ElementsByTagName(doc, "img")
  headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
```

## Deferred Function Calls

### Exercise 5.18
Without changing its behaviour, rewrite the `fetch` function to use `defer` to
close the writable file.
