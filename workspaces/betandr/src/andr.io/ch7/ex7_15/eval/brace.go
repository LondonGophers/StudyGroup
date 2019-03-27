// Add a `String` method to `Expr` to pretty-print the syntax tree. Check that the
// results, when parsed again, yield an equivalent tree.
package eval

//!+String

func (v Var) Brace() bool {
	return false
}

func (l literal) Brace() bool {
	return false
}

func (u unary) Brace() bool {
	return u.brace
}

func (b binary) Brace() bool {
	return b.brace
}

func (c call) Brace() bool {
	return c.brace
}

//!-String
