# Exercise 4.6

## Description

Write an in-place function that squashes each run of adjacent Unicode spaces in a UTF-8 encoded []byte slice into a single ASCII space.

(Unicode spaces as per Unicode's White Space property are '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP))

## Implementation
Run `go run main.go` for an example run of `SquashAdjacentSpaces` function

## Tests
Run tests with `go test`