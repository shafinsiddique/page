package main

import (
	"fmt"
	"strconv"
)

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

func getStringOfValue(value interface{}) string {
	str := ""
	switch value.(type) {
	case int:
		str =  strconv.Itoa(value.(int))
	case string:
		str = "'" + value.(string) + "'"
	case []interface{}:
		str = "(list "
		for i, v := range value.([]interface{}) {
			str += getStringOfValue(v)
			if i < len(value.([]interface{})) - 1 {
				str += " "
			}
		}
		str += ")"
	}

	return str
}

func evaluateExpr(expression IExpression) {
	//switch exprType := expression.GetType() ; exprType{
	//	case STRING_EXPR:
	//		fmt.Println("'" + expression.Evaluate().(string) + "'")
	//default:
	//	fmt.Println(expression.Evaluate())
	//}
	fmt.Println(getStringOfValue(expression.Evaluate()))
}

