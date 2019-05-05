# Exercise 1.7

The function call `io.Copy(dst, src)` reads from `src` and writes to `dst`. Use it instead of `ioutil.ReadAll` to copy
the response body to `os.Stdout` without requiring a buffer large enough to hold the entire stream. Be sure to check
the error result of `io.Copy`.
