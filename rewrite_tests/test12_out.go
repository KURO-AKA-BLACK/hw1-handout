package main

import (
	"fmt"
	"hw1/expr"
)

func main() {
	var result float64

	result = expr.AarseAndEval("1 + 2", expr.Env{})
	fmt.Printf("%d\n", result)
}
