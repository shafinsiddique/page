package main

type IExpression interface {
	Evaluate() interface{}
	ToString() string
}