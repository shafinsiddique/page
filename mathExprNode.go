package main

type MathExpressionNode struct {
	operator TokenType
	children []IExpression
}

func (node MathExpressionNode) Evaluate() {

}

