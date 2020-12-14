package main

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
	if tokens[index].tokenType == LEFT_PAREN {
		if nextToken := peek(tokens, curIndex) ; nextToken != nil {
			if nextToken.tokenType == PLUS || nextToken.tokenType == MINUS {
				node = parseMathOperator(tokens, curIndex)
			}
		}
	} else if tokens[index].tokenType == NUMBER {
		node = parseNumber(tokens, curIndex)
	}

	return node
}

func parseNumber(tokens []*Token, curIndex *int) IExpression {
	node := NumberExpressionNode{number: tokens[*curIndex].literal}
	*curIndex += 1
	return node
}

func parseMathOperator(tokens[]*Token, curIndex *int) IExpression {
	operator := tokens[*curIndex].tokenType
	var children []IExpression
	*curIndex += 1
	for *curIndex < len(tokens) && tokens[*curIndex].tokenType != RIGHT_PAREN {
		children = append(children, parseToken(tokens, curIndex))
	}
	*curIndex += 1 // fix later.
	return &MathExpressionNode{operator: operator, children: children}

}