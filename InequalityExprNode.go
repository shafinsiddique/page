package main

import "log"

type InequalityExprNode struct {
	operator TokenType
	element1 IExpression
	element2 IExpression
}

func (node InequalityExprNode) Evaluate()interface{} {
	var v1, v2 int
	if val1,ok := node.element1.Evaluate().(int) ; ! ok{
		log.Fatal(TypeMismatchError("number",val1))
	} else if val2,ok := node.element2.Evaluate().(int) ; ! ok {
		log.Fatal(TypeMismatchError("number",val2))
	} else {
		v1 = val1
		v2 = val2
	}

	if node.operator == GREATER_THAN {
		return v1 > v2
	} else if node.operator == LESS_THAN {
		return v1 < v2
	}
	return v1 == v2
}

func (node InequalityExprNode) GetType()ExpressionType {
	return INEQUALITY_EXPR
}

type AndOrExprNode struct {
	operator TokenType
	elements []IExpression
}

func hasFalse(elements []IExpression)bool {
	for _, expr := range elements {
		if val, ok := expr.Evaluate().(bool) ; ok && !val{
			return true
		} else if !ok {
			log.Fatal(TypeMismatchError("bool", val))
		}
	}
	return false
}

func hasTrue(elements []IExpression)bool {
	for _, expr := range elements {
		if val, ok := expr.Evaluate().(bool) ; ok && val {
			return true
		} else if ! ok {
			log.Fatal(TypeMismatchError("bool",val))
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

