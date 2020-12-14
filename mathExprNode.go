package main

type MathExpressionNode struct {
	operator TokenType
	children []IExpression
}

func (node MathExpressionNode) Evaluate() {

}

func (node MathExpressionNode) ToString() string {
	str := string(node.operator) + " "

	for _, expr := range node.children {
		str += expr.ToString() + " "
	}
	return str
}





