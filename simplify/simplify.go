package simplify

import "hw1/expr"
import "reflect"

func key_exist(my_key expr.Var, env expr.Env) bool {
	for key, _ := range env {
		if key == my_key {
			return true
		}
	}
	return false
}

// Depth should return the simplified expresion
func Simplify(e expr.Expr, env expr.Env) expr.Expr {
	//TODO implement the simplify
	e_type := reflect.TypeOf(e).String()
	if e_type == "expr.Var" {
		if (key_exist(e.(expr.Var), env)) {
			return expr.Expr(expr.Literal(e.Eval(env)))
		}
		return e
	} else if e_type == "expr.Unary" {
		e := e.(expr.Unary)
		e.X = Simplify(e.X, env) 
		if reflect.TypeOf(e.X).String() == "expr.Binary" {
			return e
		}
		if reflect.TypeOf(e.X).String() == "expr.Var" {
			return e
		}
		return expr.Expr(expr.Literal(e.Eval(env)))
	} else if e_type == "expr.Binary" {
		e := e.(expr.Binary)
		e.X = Simplify(e.X, env)
		e.Y = Simplify(e.Y, env)
		if reflect.TypeOf(e.X).String() == "expr.Var" && reflect.TypeOf(e.Y).String() == "expr.Var" {
			return e
		} else if reflect.TypeOf(e.X).String() == "expr.Var" {
			if (e.Y).Eval(env) == 0 {
				if e.Op == '+' {
					return e.X
				} else if e.Op == '*' {
					return e.Y
				}
			} else if (e.Y).Eval(env) == 1 {
				if e.Op == '*' {
					return e.X
				}
			} else {
				return e
			}
		} else if reflect.TypeOf(e.Y).String() == "expr.Var" {
			if (e.X).Eval(env) == 0 {
				if e.Op == '+' {
					return e.Y
				} else if e.Op == '*' {
					return e.X
				}
			} else if (e.X).Eval(env) == 1 {
				if e.Op == '*' {
					return e.Y
				}
			} else {
				return e
			}
		}
		if reflect.TypeOf(e.Y).String() == "expr.Binary" || reflect.TypeOf(e.X).String() == "expr.Binary" {
			return e
		}
		if reflect.TypeOf(e.Y).String() == "expr.Unary" || reflect.TypeOf(e.X).String() == "expr.Unary" {
			return e
		}
		return expr.Expr(expr.Literal(e.Eval(env)))
	} else if e_type == "expr.Literal" {
		return e
	} else {
		panic("panicing")
	}
}
