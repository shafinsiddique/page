package main

type NumberExpressionNode struct {
	number string
}

func (node NumberExpressionNode) Evaluate() {

}

func (node NumberExpressionNode) ToString() string {
	return node.number
}