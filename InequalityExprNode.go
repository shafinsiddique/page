package main

type InequalityExprNode struct {
	operator TokenType
	element1 int
	element2 int
}

func (node InequalityExprNode) Evaluate()interface{} {
	if node.operator == GREATER_THAN {
		return node.element1 > node.element2
	} else if node.operator == LESS_THAN {
		return node.element1 < node.element2
	}

	return node.element2 == node.element1
}

func (node InequalityExprNode) GetType()ExpressionType {
	return INEQUALITY_EXPR
}
