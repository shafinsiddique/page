package main

type IParser interface {
	Parse(tokens []Token)*AST
}
