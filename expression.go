package main

type IExpression interface {
	Evaluate()
	ToString() string
}