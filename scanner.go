package main

type IScanner interface {
	GetTokens() []Token
}
