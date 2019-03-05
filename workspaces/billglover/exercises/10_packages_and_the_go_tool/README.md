# Chapter 10 - Packages and the Go Tool

## Blank Imports

### Exercise 10.1
Extend the `jpeg` program so that it converts any supported input format to any
output format, using `image.Decode` to detect the input format and a flag to
select the output format.

### Exercise 10.2
Define a generic archive file-reading function capable of reading ZIP files
(`archive/zip`) and POSIX tar files (`archive/tar`). Use a registration
mechanism similar to the one described above so that support for each file
format can be plugged in using blank imports.

## Downloading Packages

### Exercise 10.3
Using `fetch http://gopl.io/ch1/helloworld?go-get=1`, find out which service
hosts the code samples for this book. (HTTP reqests from `go get` include the
`go-get` parameter so that servers can distinguish them from ordinary browser
requests.)

## Querying Packages

### Exercise 10.4
Construct a tool that reports the set of all packages in the workspace that
transitively depend on the packages specified by the arguments. Hint: you will
need to run `go list` twice, once for the initial packages and once for all
packages. You may want to parse its JSON output using the `encoding/json`
package (§4.5).
