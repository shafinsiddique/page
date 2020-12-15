package main

import (
	"log"
	"strconv"
	"unicode"
)

type BasicScanner struct {
}

func isDigit(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

func NewBasicScanner() IScanner {
	return BasicScanner{}
}

func (scanner BasicScanner) GetTokens(strScan string) []*Token{
	var tokens []*Token
	str := strScan
	starting := 0
	curIndex := &starting
	for *curIndex < len(str) {
		index := *curIndex
		incremented := false
		curChar := string(str[index])
		var token *Token
		if val,ok := SINGLE_CHAR_TOKENS[curChar] ; ok {
			token = &Token{TokenType: val.TokenType}
		} else if isDigit(curChar) {
			token = scanNumberToken(&str, curIndex)
			incremented = true
		} else if curChar == "\"" {
			token = scanStringToken(&str, curIndex)
			incremented = true
		} else {
			cache := *curIndex
			word := tryGetWord(&str, curIndex)
			if t, ok := RESERVED_WORD_TOKENS[word] ; ok {
				token = &Token{TokenType: t.TokenType}
				incremented = true
			} else {
				*curIndex = cache
			}
		}

		if token != nil {
			tokens = append(tokens, token)
		}

		if !incremented {
			*curIndex += 1
		}
	}

	return tokens
}

func tryGetWord(str *string, index *int) string {
	word := ""
	for *index < len(*str) {
		if unicode.IsLetter(rune((*str)[*index])) {
			word += string((*str)[*index])
		} else {
			break
		}
		*index += 1
	}

	return word
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

	return &Token{TokenType: NUMBER, Literal: num}
}

func scanStringToken(str *string, index *int) *Token {
	strVal := ""
	*index += 1 // we don't need to count the first quote.
	for *index < len(*str) {
		if curChar := string((*str)[*index]) ; curChar != "\"" {
			strVal += curChar
		} else {
			break
		}
		*index += 1
	}

	if *index == len(*str) {
		log.Fatal(UNCLOSED_STRING)
	}

	*index += 1
	return &Token{TokenType: STRING, Literal: strVal}
}

