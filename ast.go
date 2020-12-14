package main

import "fmt"

type AST struct {
	expressions []IExpression
}

func (ast AST) ToString() string {
	str := ""
	for _, expr := range ast.expressions {
		str += expr.ToString() + "\n"
	}

	return str
}

func (ast AST) Evaluate() {
	for _, expr := range ast.expressions {
		fmt.Println(expr.Evaluate())
	}
}