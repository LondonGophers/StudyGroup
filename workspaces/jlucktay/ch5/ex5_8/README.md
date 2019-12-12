# Exercise 5.8

Modify `forEachNode` so that the `pre` and `post` functions return a boolean result indicating whether to continue the
traversal. Use it to write a function `ElementByID` with the following signature that finds the first HTML element with
the specified `id` attribute. The function should stop the traversal as soon as a match is found.

``` go
    func ElementByID(doc *html.Node, id string) *html.Node
```
