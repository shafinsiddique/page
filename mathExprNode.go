package main

import "log"

type MathExpressionNode struct {
	operator TokenType
	children []IExpression
}

func (node MathExpressionNode) Evaluate() interface{} {
	index := 0
	value := 0

	for index < len(node.children) {
		if val, ok := node.children[index].Evaluate().(int); ! ok {
			log.Fatal(TypeMismatchErrorNumber(val))
		} else if node.operator == PLUS {
			value += val
		} else if node.operator == MINUS {
			value -= val
		} else if node.operator == DIVIDE {
			value /= val
		} else {
			value *= val
		}

		index += 1
	}

	return value
}


func (node MathExpressionNode) GetType()ExpressionType {
	return MATH_EXPR
}
