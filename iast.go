package main

type IAST interface {
	AddExpression(expr IExpression)
	GetExpressions()[]IExpression
}


