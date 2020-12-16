package main

import "log"

type IfExprNode struct {
	condition IExpression
	thenExp IExpression
	elseExp IExpression
}

func (node IfExprNode) Evaluate()interface{} {
	var result interface{}
	if cond, isBool := node.condition.Evaluate().(bool) ; !isBool {
		log.Fatal(TypeMismatchError("bool", cond))
	} else {
		if cond {
			result = node.thenExp.Evaluate()
		} else {
			result = node.elseExp.Evaluate()
		}
	}
	return result
}

func (node IfExprNode) GetType()ExpressionType {
	return IF_EXPR
}
