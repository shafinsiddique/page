package main

type IAST interface {
	ToString() string
	Evaluate()
	AddExpression(expr IExpression)
}


