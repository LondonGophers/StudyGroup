// Add a `String` method to `Expr` to pretty-print the syntax tree. Check that the
// results, when parsed again, yield an equivalent tree.
package eval

//!+String

func (v Var) String() string {
	return ""
}

func (literal) String() string {
	return ""
}

func (u unary) String() string {
	return ""
}

func (b binary) String() string {
	return ""
}

func (c call) String() string {
	return ""
}

//!-String
