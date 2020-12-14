package main

import "fmt"

type Evaluator struct {
}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (evaluator Evaluator) Evaluate(ast IAST) {
	for _, expr := range ast.GetExpressions() {
		fmt.Println(expr.Evaluate())
	}
}

func (evaluator Evaluator) EvaluateLast(ast IAST) {
	expressions := ast.GetExpressions()
	if len(expressions) > 0 {
		fmt.Println(expressions[len(expressions)-1].Evaluate())
	}
}
