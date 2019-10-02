# Chapter 12 - Reflection

## Display, a Recursive Value Printer

### Exercise 12.1
Extend `Display` so that it can display maps whoes kets are structs or arrays.

### Exercise 12.2
Make `display` safe to use on cyclic data structures by bounding the number of
steps it takes before abandoning the recursion. (In Section 13.3, we'll see
another way to detect cycles.)

## Example: Encoding S-Expressions

### Exercise 12.3
Implement the missing cases of the `encode` function. Encode booleans as `t` and
`nil`, floating-point numbers using Go's notation, and complex numbers like
1+2 _i_ as `#C1.0 2.0)`. Interfaces can be encoded as a pair of a type name and
a value, for instance `("[]int" (1 2 3))`, but beware that this notation is
ambiguous: the `reflect.Type.String` method may return the same string for
different types.

### Exercise 12.4
Modify `encode` to pretty-print the S-expression in the style shown above.

### Exercise 12.5
Adapt `encode` to emit JSON instead of S-expressions. Test your encoder using
the standard decoder, `json.Unmarshal`.

### Exercise 12.6
Adapt `encode` so that, as an optimization, it does not encode a field whose
value is the zero value of its type.

### Exercise 12.7
Create a streaming API for the S-expressions encoder, following the style of
`json.Encoder` (§4.5).

## Example: Decoding S-Expressions

### Exercise 12.8
The `sexpr.Unmarshal` function, like `json.Unmarshal`, requires the complete
input in a byte slice before it can begin decoding. Define a `sexpr.Decoder`
type that, like `json.Decoder`, allows a sequence of values to be decoded from
an `io.Reader`. Change `sexpr.Unmarshal` to use this new type.

### Exercise 12.9
Write a token-based API for decoding S-expressions, following the style of
`xml.Decoder` (§7.14). You will need 5 types of tokens: `Symbol`, `String`,
`Int`, `StartList`, and `EndList`.

### Exercise 12.10
Extend `sexpr.Unmarshal` to handle the booleans, floating-point numbers, and
interfaces encoded by your solution to Exercise 12.3. (Hint: to decode
interfaces, you will need a mapping from the name of each supported type to
its `reflect.Type`.)

## Accessing Struct Field Tags

### Exercise 12.11
Write the corresponding `Pack` function. Given a struct value, `Pack` should
return a URL incorporating the parameter values from the struct.

### Exercise 12.12
Extend the field tag notation to express parameter validity requirements. For
example, a string might need to be a valid email address or credit-card number,
and an integer might need to be a valid US ZIP code. Modify `Unpack` to check
these requirements.

### Exercise 12.13
Modify the S-expression encoder (§12.4) and decoder (§12.6) so that they honor
the `sexpr:"...` field tag in a similar manner to `encoding/json` (§4.5).
