package main

import "fmt"

type Evaluator struct {
}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (evaluator Evaluator) Evaluate(ast IAST) {
	for _, expr := range ast.GetExpressions() {
		evaluateExpr(expr)
	}
}

func (evaluator Evaluator) EvaluateLast(ast IAST) {
	expressions := ast.GetExpressions()
	if len(expressions) > 0 {
		evaluateExpr(expressions[len(expressions)-1])
	}
}

func evaluateExpr(expression IExpression) {
	switch exprType := expression.GetType() ; exprType{
		case STRING_EXPR:
			fmt.Println("'" + expression.Evaluate().(string) + "'")
	default:
		fmt.Println(expression.Evaluate())
	}
}

