// Define a new concrete type that satisfies the `Expr` interface and provides a
// new operation such as computing the minimum value of its operands. Since the
// `Parse` function does not create instances of this new type, to use it you will
// need to construct a syntax tree directly (or extend the parser).
package eval

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
	// Pretty-print the syntax tree
	String() string
	// Brace() reports if this Expr has an enclosing brace
	Brace() bool
}

//!+ast

// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op    rune // one of '+', '-'
	x     Expr
	brace bool
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op    rune // one of '+', '-', '*', '/'
	x, y  Expr
	brace bool
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn    string // one of "pow", "sin", "sqrt"
	args  []Expr
	brace bool
}

//!-ast
