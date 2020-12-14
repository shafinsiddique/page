package main

type IParser interface {
	Parse(ast IAST, tokens []*Token)
}
