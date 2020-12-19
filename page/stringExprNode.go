package main

type StringExprNode struct {
	stringLiteral string
}

func (node StringExprNode) Evaluate()interface{} {
	return node.stringLiteral
}

func (node StringExprNode) GetType()ExpressionType {
	return STRING_EXPR
}
