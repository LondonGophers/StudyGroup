package eval

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"min(x, y)", Env{"x": 3, "y": 4}, "3"},
		{"min(y, x)", Env{"x": 5, "y": 4}, "4"},
		{"min(x, x)", Env{"x": 0}, "0"},
		{"min(a, b)", Env{"a": 1.2, "b": 3.4}, "1.2"},
	}
	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
