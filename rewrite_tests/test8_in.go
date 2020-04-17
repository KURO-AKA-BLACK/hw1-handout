package main

import (
	"fmt"
)

func main() {
	var result float64

	result = 2 + expr.ParseAndEval("1 + 2", expr.Env{})
	fmt.Printf("%d\n", result)
}
