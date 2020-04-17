package rewrite

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"reflect"
	"hw1/expr"
	"hw1/simplify"
	"strconv" // This package may be helpful...
)
func make_string(input expr.Expr) string{
	var my_string string
	if reflect.TypeOf(input).String() == "expr.Var"{
		temp := input.(expr.Var)
		return my_string + string(temp)
	} else if reflect.TypeOf(input).String() == "expr.Literal" {
		temp := float64(input.(expr.Literal))
		return my_string + strconv.FormatFloat(temp, 'f', -1, 64)
	} else if reflect.TypeOf(input).String() == "expr.Unary" {
		temp := input.(expr.Unary)
		return my_string + string(temp.Op) + make_string(temp.X)
	} else {
		temp := input.(expr.Binary)
		return my_string + make_string(temp.X) + " " + string(temp.Op) + " " + make_string(temp.Y)
	}
}
// rewriteCalls should modify the passed AST
var g_var int = 0;
var im int = 0;

func rewriteCalls(node ast.Node) {
	//TODO Write the rewriteCalls function
	g_var = 0
	im = 0
	ast.Inspect(node, func(n ast.Node) bool {
		switch n.(type){
			case *ast.ImportSpec:
				this_spec := n.(*ast.ImportSpec)
				if reflect.TypeOf(this_spec.Path).String() == "*ast.BasicLit" {
					this_path := this_spec.Path
					if this_path.Value == "\"hw1/expr\"" {
						im = 1;
					}
				}
			case *ast.FuncDecl:
				if g_var == 1 || im == 0 {
					return true
				}
				first := n.(*ast.FuncDecl)
				if reflect.TypeOf(first.Name).String() == "*ast.Ident" {
					sss := first.Name
					if sss.Name == "ParseAndEval" {
						g_var = 1
						return true
					}
				}
			case *ast.CallExpr:
				if g_var == 1 || im == 0 {
					return true
				}
				this_func := (n.(*ast.CallExpr)).Fun
				if reflect.TypeOf(this_func).String() == "*ast.SelectorExpr" {
					this_x := (this_func.(*ast.SelectorExpr)).X
					this_sel := (this_func.(*ast.SelectorExpr)).Sel
					if reflect.TypeOf(this_x).String() == "*ast.Ident" {
						if (this_x.(*ast.Ident)).Name != "expr" {
							return true
						}
					}
					if reflect.TypeOf(this_sel).String() == "*ast.Ident" {
						if this_sel.Name != "ParseAndEval" {
							return true
						}
					}
				}
				this_args := (n.(*ast.CallExpr)).Args
				if len(this_args) != 2 {
					return true
				}
				if reflect.TypeOf(this_args).String() == "[]ast.Expr" {
					this_zero := this_args[0]
					if reflect.TypeOf(this_zero).String() == "*ast.BasicLit" {
						this_value := (this_zero.(*ast.BasicLit)).Value
						this_value = string([]rune(this_value)[1:len(this_value)-1])
						this_e, err := expr.Parse(this_value)
						if err != nil {
							return true
						}
						var this_env expr.Env
						this_out := simplify.Simplify(this_e, this_env)
						my_string := make_string(this_out)
						my_string = "\"" + my_string + "\""
						addr_1 := &(((*(n.(*ast.CallExpr))).Args[0]).(*ast.BasicLit)).Value
						*addr_1 = my_string
					}
				}
			default: 
			
		}
		return true
	})
	//panic("TODO: implement this!")
}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}
	
	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
