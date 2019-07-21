// Add a `String` method to `Expr` to pretty-print the syntax tree. Check that the
// results, when parsed again, yield an equivalent tree.
package eval

import (
	"fmt"
	"testing"
)

// todo check difference between parsing the non-parenthesis and parenthesis versions

func TestString(t *testing.T) {
	tests := []string{
		"x * y",
		"x / x",
		"y - x",
		"sin(x)",
		"-x",
		"+y",
		"sqrt(A / pi)",
		"pow(x, 3) + pow(y, 3)",
		"2 * 2 + 2",
		"2 * (2 + 2)",
		"5 / 9 * (F - 32)",
	}
	for _, expr := range tests {
		fmt.Printf("\n%s\n", expr)

		e, _ := Parse(expr)
		got := e.String()

		if got != expr {
			t.Errorf("want \"%s\" got \"%s\"\n", expr, got)
		}
	}
}
