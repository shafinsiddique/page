package main

import "strconv"

type NumberExpressionNode struct {
	number string
}

func (node NumberExpressionNode) Evaluate() interface{} {
	v, _ := strconv.Atoi(node.number)
	return v
}

func (node NumberExpressionNode) GetType() ExpressionType {
	return NUMBER_EXPR
}
