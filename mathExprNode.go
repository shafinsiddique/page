package main

type MathExpressionNode struct {
	operator TokenType
	children []IExpression
}

func (node MathExpressionNode) Evaluate() interface{} {
	cur_value := node.children[0].Evaluate().(int)
	index := 1
	for index < len(node.children) {
		value := node.children[index].Evaluate().(int)
		if node.operator == PLUS {
			cur_value += value
		} else if node.operator == MINUS {
			cur_value -= value
		} else if node.operator == DIVIDE {
			cur_value /= value
		} else {
			cur_value *= value
		}

		index += 1
	}

	return cur_value
}

func (node MathExpressionNode) ToString() string {
	str := string(node.operator) + " "

	for _, expr := range node.children {
		str += expr.ToString() + " "
	}
	return str
}





