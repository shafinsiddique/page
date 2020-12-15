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

type AndOrExprNode struct {
	operator TokenType
	elements []bool
}

func hasFalse(elements []bool)bool {

	for _, v := range elements {
		if !v {
			return true
		}
	}
	return false
}

func hasTrue(elements []bool)bool {
	for _, v := range elements {
		if v {
			return true
		}
	}
	return false
}

func(node AndOrExprNode) Evaluate()interface{} {
	if node.operator == AND {
		return !hasFalse(node.elements)
	}
	return hasTrue(node.elements)
}

func (node AndOrExprNode) GetType()ExpressionType {
	return AND_OR_EXPR
}

