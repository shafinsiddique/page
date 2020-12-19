package main


type IExpression interface {
	Evaluate() interface{}
	GetType() ExpressionType
}

