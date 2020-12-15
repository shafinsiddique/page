package main

type IAST interface {
	//ToString() string
	AddExpression(expr IExpression)
	GetExpressions()[]IExpression
}


