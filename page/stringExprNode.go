package main

import "log"

type StringExprNode struct {
	stringLiteral string
}

func (node StringExprNode) Evaluate()interface{} {
	return node.stringLiteral
}

func (node StringExprNode) GetType()ExpressionType {
	return STRING_EXPR
}

type StringEqualsExprNode struct {
	strs []IExpression
}

func (node StringEqualsExprNode) Evaluate()interface{} {
	if val1, ok := node.strs[0].Evaluate().(string) ; ok {
		if val2, ok := node.strs[1].Evaluate().(string) ; ok {
			return val1 == val2
		} else {
			log.Fatal(TypeMismatchError("string", val2))
		}
	} else {
		log.Fatal(TypeMismatchError("string", val1))

	}
	return nil

}

func (node StringEqualsExprNode) GetType()ExpressionType {
	return STRING_EXPR
}
