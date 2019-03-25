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
		{"min(x, y, z)", Env{"x": 3, "y": 4, "z": 5}, "3"},
		{"min(z, y, x)", Env{"x": 5, "y": 4, "z": 3}, "3"},
		{"min(x)", Env{"x": 0}, "0"},
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
