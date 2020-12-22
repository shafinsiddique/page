package main


type Token struct {
	TokenType TokenType
	Literal   string
}

func (t Token) ToString() string {
	return "TokenType: " + string(t.TokenType) + "\n" + "Literal: " + t.Literal + "\n"
}

var SINGLE_CHAR_TOKENS = map[string]*Token {
	"+": {TokenType: PLUS},
	"-":{TokenType: MINUS},
	"*":{TokenType: TIMES},
	"/":{TokenType: DIVIDE},
	"(":{TokenType: LEFT_PAREN},
	")":{TokenType: RIGHT_PAREN},
	">":{TokenType: GREATER_THAN},
	"<":{TokenType: LESS_THAN},
	"=":{TokenType: EQUAL},
	"%":{TokenType: MOD},
}

var PEEK_MAP = map[string]string{
	">":"=",
	"<":"=",
}

var TWO_CHAR_TOkENS = map[string]*Token {
	">=":{TokenType: GREATER_THAN_EQUAL},
	"<=":{TokenType: LESS_THAN_EQUAL},
}
var RESERVED_WORD_TOKENS = map[string]*Token {
	"list":{TokenType: LIST},
	"cons":{TokenType: CONS},
	"car":{TokenType: FIRST},
	"and":{TokenType: AND},
	"or":{TokenType: OR},
	"True":{TokenType: TRUE},
	"False":{TokenType: FALSE},
	"if":{TokenType: IF},
	"define":{TokenType: DEFINE},
	"cdr":{TokenType: CDR},
	"length":{TokenType: LENGTH},
	"equals":{TokenType: EQUALS },
}
