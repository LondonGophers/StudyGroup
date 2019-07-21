// Add a `String` method to `Expr` to pretty-print the syntax tree. Check that the
// results, when parsed again, yield an equivalent tree.
package eval

import (
	"fmt"
	"strconv"
	"strings"
)

//!+String

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', -1, 64)
}

func (u unary) String() string {
	return fmt.Sprintf("%s%s", string(u.op), u.x)
}

func (b binary) String() string {
	var sb strings.Builder
	if b.Brace() {
		sb.WriteString("(")
	}
	sb.WriteString(fmt.Sprintf("%s %s %s", b.x, string(b.op), b.y))
	if b.Brace() {
		sb.WriteString(")")
	}

	return sb.String()
}

func (c call) String() string {
	var separator string

	if len(c.args) > 1 {
		separator = ", "
	}

	var args strings.Builder
	for i, arg := range c.args {
		args.WriteString(arg.String())
		if i < len(c.args)-1 {
			args.WriteString(separator)
		}
	}

	return fmt.Sprintf("%s(%s)", c.fn, args.String())
}

//!-String
