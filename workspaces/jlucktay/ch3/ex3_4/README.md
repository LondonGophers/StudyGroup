# Exercise 3.4

Following the approach of the Lissajous example in Section 1.7, construct a web server that computes surfaces and
writes SVG data to the client. The server must set the `Content-Type` header like this:

``` go
w.Header().Set("Content-Type", "image/svg+xml")
```

(This step was not required in the Lissajous example because the server uses standard heuristics to recognize common
formats like PNG from the first 512 bytes of the response, and generates the proper header.) Allow the client to
specify values like height, width, and color as HTTP request parameters.
