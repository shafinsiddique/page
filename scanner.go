package main

type IScanner interface {
	GetTokens(strToScan string) []*Token
}
