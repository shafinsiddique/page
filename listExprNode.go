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

type FirstExpressionNode struct {
	list []interface{}
}

func (node FirstExpressionNode) Evaluate()interface{} {
	if len(node.list) < 1 {
		log.Fatal(FIRST_ON_EMPTY)
	}
	return node.list[0]
}

func (node FirstExpressionNode) GetType() ExpressionType {
	return FIRST_EXPR
}