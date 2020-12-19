package main

type BooleanExprNode struct {
	operator TokenType
}

func (node BooleanExprNode) Evaluate()interface{} {
	return node.operator == TRUE
}

func (node BooleanExprNode) GetType()ExpressionType {
	return BOOLEAN_EXPR
}