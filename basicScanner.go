package main

import (
	"strconv"
)

type BasicScanner struct {
	strScan string
	curIndex int
}

func isDigit(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}
func (scanner *BasicScanner) GetTokens() []*Token{
	var tokens []*Token
	str := scanner.strScan
	curIndex := &scanner.curIndex
	for *curIndex < len(str) {
		index := *curIndex
		incremented := false
		if curChar := string(str[index]) ; curChar == "(" {
			tokens = append(tokens, &Token{tokenType: LEFT_PAREN, literal: "("})
		} else if curChar == ")" {
			tokens = append(tokens, &Token{tokenType: RIGHT_PAREN, literal: ")"})
		} else if curChar == "+" {
			tokens = append(tokens, &Token{tokenType: PLUS, literal: "+"})
		} else if curChar == "-" {
			tokens = append(tokens, &Token{tokenType: MINUS, literal:"-"})
		} else if isDigit(curChar) {
			tokens = append(tokens, scanNumberToken(&str, curIndex))
			incremented = true
		}
		if !incremented {
			*curIndex += 1
		}
	}
	return tokens
}

func scanNumberToken(str *string, index *int) *Token{
	curStr := *str
	num := ""
	for *index < len(curStr) {
		if curChar := string(curStr[*index]) ; isDigit(curChar){
			num += curChar
			*index += 1
		} else {
			break
		}
	}
	return &Token{tokenType: NUMBER, literal: num}
}

