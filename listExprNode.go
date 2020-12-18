package main

import "log"

type ListExpressionNode struct {
	elements []IExpression
}

func (node ListExpressionNode) Evaluate()interface{} {
	var elements []interface{}
	for _, expr := range node.elements {
		elements = append(elements, expr.Evaluate())
	}
	return elements
}

func (node ListExpressionNode) GetType() ExpressionType {
	return LIST_EXPR
}

type ConsExpressionNode struct {
	element interface{}
	list []interface{}
}

func (node ConsExpressionNode) Evaluate()interface{} {
	var elements []interface{}
	elements = append(elements, node.element)

	for _, expr := range node.list {
		elements = append(elements, expr)
	}
	return elements
}

func (node ConsExpressionNode) GetType()ExpressionType {
	return CONS_EXPR
}

type CarCdrExpressionNode struct {
	tokenType TokenType
	list      IExpression
}

func (node CarCdrExpressionNode) Evaluate()interface{} {
	var result interface{}
	if val, ok := node.list.Evaluate().([]interface{}); ok {
		if len(val) > 0 {
			if node.tokenType == CDR  {
				result = val[1:]
			} else {
				result = val[0]
			}
		} else {
			log.Fatal(LIST_FUNCTION_ON_EMPTY)
		}
	} else {
		log.Fatal(TypeMismatchError("list", val))
	}

	return result
}

func (node CarCdrExpressionNode) GetType() ExpressionType {
	return FIRST_EXPR
}