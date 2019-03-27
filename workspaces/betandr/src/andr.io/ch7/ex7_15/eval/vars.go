// Copyright © 2019 Beth Anderson
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

//!+Vars

func (v Var) Vars() []Var {
	return []Var{v}
}

func (l literal) Vars() []Var {
	return nil
}

func (u unary) Vars() []Var {
	return nil
}

func (b binary) Vars() []Var {
	return append(b.x.Vars(), b.y.Vars()...)
}

func (c call) Vars() []Var {
	args := make([]Var, 0)

	for _, arg := range c.args {
		args = append(args, arg.Vars()...)
	}

	return args
}

//!-Vars
