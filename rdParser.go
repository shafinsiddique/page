package main

import "log"

type RDParser struct {
}

func NewRDParser() IParser {
	return RDParser{}
}

func (p RDParser) Parse(ast IAST, tokens []*Token) {
	index := 0
	for index < len(tokens) {
		ast.AddExpression(parseToken(tokens, &index))
	}
}

func peek(tokens[]*Token, index *int)  *Token {
	var token *Token
	if *index + 1 < len(tokens){
		token = tokens[*index+1]
	}
	*index += 1
	return token
}

func parseToken(tokens []*Token, curIndex *int) IExpression {
	index := *curIndex
	var node IExpression
	if tokens[index].TokenType == LEFT_PAREN {
		if nextToken := peek(tokens, curIndex) ; nextToken != nil {
			if nextToken.TokenType == PLUS || nextToken.TokenType == MINUS ||
				nextToken.TokenType == DIVIDE || nextToken.TokenType == TIMES{
				node = parseMathOperator(tokens, curIndex)
			}
		} else {
			log.Fatal(UNCLOSED_PARENTHESIS)
		}
	} else if tokens[index].TokenType == NUMBER {
		node = parseNumber(tokens, curIndex)
	}

	return node
}

func parseNumber(tokens []*Token, curIndex *int) IExpression {
	node := NumberExpressionNode{number: tokens[*curIndex].Literal}
	*curIndex += 1
	return node
}

func parseMathOperator(tokens[]*Token, curIndex *int) IExpression {
	operator := tokens[*curIndex].TokenType
	var children []IExpression
	*curIndex += 1
	for *curIndex < len(tokens) && tokens[*curIndex].TokenType != RIGHT_PAREN {
		children = append(children, parseToken(tokens, curIndex))
	}
	if *curIndex == len(tokens) {
		// this means we went to the end of the tokens without seeing a closing parenthesis.
		log.Fatal(UNCLOSED_PARENTHESIS)
	}

	if len(children) < 2 {
		log.Fatal(TOO_FEW_ARGUMENTS)
	}

	*curIndex += 1
	return &MathExpressionNode{operator: operator, children: children}

}