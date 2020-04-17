package depth

import "hw1/expr"
import "reflect"
import "math"

// Depth returns the maximum number of AST nodes between the
// root of the tree and any leaf (literal or variable) in the tree.
func Depth(e expr.Expr) uint {
	// TODO: implement this function
	e_type := reflect.TypeOf(e).String()
	if e_type == "expr.Literal" || e_type == "expr.Var" {
		return 1
	} else if e_type == "expr.Unary" {
		e := e.(expr.Unary)
		return 1 + Depth(e.X)
	} else if e_type == "expr.Binary" {
		e := e.(expr.Binary)
		return uint(math.Max(float64(1 + Depth(e.X)), float64(1 + Depth(e.Y))))
	} else {
		panic("Something's wrong")
	}
}
