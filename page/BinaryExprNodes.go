package main

import "log"

type BinaryExprNodes struct {
	operator TokenType
	first IExpression
	second IExpression
}

func (node BinaryExprNodes) Evaluate()interface{} {
	val1, ok := node.first.Evaluate().(int)
	if !ok {
		log.Fatal(TypeMismatchError("number", val1))
	}
	val2, ok := node.second.Evaluate().(int)

	if ! ok {
		log.Fatal(TypeMismatchError("number", val2))
	}

	return val1 % val2
}

func (node BinaryExprNodes) GetType() ExpressionType {
	return BOOLEAN_EXPR
}

