package main

import (
	"log"
)

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

func peek(tokens []*Token, index *int) *Token {
	var token *Token
	if *index+1 < len(tokens) {
		token = tokens[*index+1]
	}
	*index += 1
	return token
}

func parseToken(tokens []*Token, curIndex *int) IExpression {
	index := *curIndex
	var node IExpression
	if tokens[index].TokenType == LEFT_PAREN {
		if nextToken := peek(tokens, curIndex); nextToken != nil {
			if nextToken.TokenType == PLUS || nextToken.TokenType == MINUS ||
				nextToken.TokenType == DIVIDE || nextToken.TokenType == TIMES {
				node = parseMathOperator(tokens, curIndex)
			} else if nextToken.TokenType == LIST {
				node = parseList(tokens, curIndex)
			} else if nextToken.TokenType == CONS {
				node = parseCons(tokens, curIndex)
			} else if nextToken.TokenType == FIRST {
				node = parseFirst(tokens, curIndex)
			} else if nextToken.TokenType == GREATER_THAN || nextToken.TokenType == LESS_THAN ||
				nextToken.TokenType == EQUAL  || nextToken.TokenType == GREATER_THAN_EQUAL || nextToken.TokenType == LESS_THAN_EQUAL{
				node = parseInequality(tokens ,curIndex)
			} else if nextToken.TokenType == AND || nextToken.TokenType == OR {
				node = parseAndOr(tokens, curIndex)
			} else if nextToken.TokenType == IF {
				node = parseIf(tokens, curIndex)
			}
		} else {
			log.Fatal(UNCLOSED_PARENTHESIS)
		}
	} else if tokens[index].TokenType == NUMBER {
		node = parseNumber(tokens, curIndex)
	} else if tokens[index].TokenType == STRING {
		node = parseString(tokens, curIndex)
	} else if tokens[index].TokenType == TRUE || tokens[index].TokenType == FALSE {
		node = BooleanExprNode{operator: tokens[index].TokenType}
		*curIndex += 1
	} else {
		log.Fatal(UNEXPECTED_TOKEN)
	}
	return node
}

func parseIf(tokens[]*Token, curIndex *int)IExpression {
	*curIndex += 1
	args := getAllArguments(tokens, curIndex)
	checkIfCorrectArguments(3, len(args))
	*curIndex += 1
	return IfExprNode{condition: args[0], thenExp: args[1], elseExp: args[2]}
}

func getAllArguments(tokens[]*Token, curIndex *int)[]IExpression {
	elements := []IExpression{}
	for *curIndex < len(tokens) && tokens[*curIndex].TokenType != RIGHT_PAREN{
		elements = append(elements, parseToken(tokens, curIndex))
	}
	if *curIndex == len(tokens) {
		log.Fatal(UNCLOSED_PARENTHESIS)
	}

	return elements

}
func parseAndOr(tokens []*Token, curIndex *int)IExpression {
	curToken := tokens[*curIndex]
	*curIndex += 1
	elements := getAllArguments(tokens, curIndex)
	if len(elements) < 2 {
		log.Fatal(TOO_FEW_ARGUMENTS)
	}

	*curIndex += 1
	return AndOrExprNode{operator: curToken.TokenType, elements: elements}

}
func checkIfCorrectArguments(expected int, actual int) {
	if actual < expected {
		log.Fatal(TOO_FEW_ARGUMENTS)
	} else if actual > expected {
		log.Fatal(TOO_MANY_ARGUMENTS)
	}
}

func parseInequality(tokens []*Token, curIndex *int)IExpression {
	elements := []IExpression{}
	curToken := tokens[*curIndex]
	*curIndex += 1

	for *curIndex < len(tokens) && tokens[*curIndex].TokenType != RIGHT_PAREN {
		elements = append(elements, parseToken(tokens, curIndex))
	}

	if *curIndex == len(tokens) {
		log.Fatal(UNCLOSED_PARENTHESIS)
	}
	checkIfCorrectArguments(2, len(elements))
	// if it comes here, it means it has 2.
	*curIndex += 1
	return InequalityExprNode{element1: elements[0], element2: elements[1], operator: curToken.TokenType}
}

func parseFirst(tokens []*Token, curIndex *int) IExpression {
	elements := []IExpression{}
	*curIndex += 1
	for *curIndex < len(tokens) && tokens[*curIndex].TokenType != RIGHT_PAREN {
		elements = append(elements, parseToken(tokens, curIndex))
	}

	if *curIndex == len(tokens) {
		log.Fatal(UNCLOSED_PARENTHESIS)
	}

	if len(elements) < 1 {
		log.Fatal(TOO_FEW_ARGUMENTS)
	} else if len(elements) > 1 {
		log.Fatal(TOO_MANY_ARGUMENTS)
	}

	if elements[0].GetType() != LIST_EXPR {
		log.Fatal(TypeMismatcErrorList(elements[0].GetType()))
	}

	*curIndex += 1

	return FirstExpressionNode{list:elements[0].Evaluate().([]interface{})}

}
func parseCons(tokens []*Token, curIndex *int) IExpression {
	elements := []IExpression{}
	*curIndex += 1
	for *curIndex < len(tokens) && tokens[*curIndex].TokenType != RIGHT_PAREN {
		elements = append(elements, parseToken(tokens, curIndex))
	}

	if *curIndex == len(tokens) {
		log.Fatal(UNCLOSED_PARENTHESIS)
	}

	if len(elements) < 2 {
		log.Fatal(TOO_FEW_ARGUMENTS)
	} else if len(elements) > 2 {
		log.Fatal(TOO_MANY_ARGUMENTS)
	}

	if elements[1].GetType() != LIST_EXPR {
		log.Fatal(TypeMismatcErrorList(elements[1].GetType()))
	}
	*curIndex += 1
	return ConsExpressionNode{element: elements[0].Evaluate(), list:elements[1].Evaluate().([]interface{})}
}
func parseList(tokens []*Token, curIndex *int) IExpression {
	elements := []IExpression{}
	*curIndex += 1
	for *curIndex < len(tokens) && tokens[*curIndex].TokenType != RIGHT_PAREN {
		elements = append(elements, parseToken(tokens, curIndex))
	}
	if *curIndex == len(tokens) {
		// no right paren found.
		log.Fatal(UNCLOSED_PARENTHESIS)
	}

	*curIndex += 1 // skipping the right paren, since it exists.
	node := ListExpressionNode{elements: elements}
	return node
}

func parseString(tokens []*Token, curIndex *int) IExpression {
	node := StringExprNode{stringLiteral: tokens[*curIndex].Literal}
	*curIndex += 1
	return node
}

func parseNumber(tokens []*Token, curIndex *int) IExpression {
	node := NumberExpressionNode{number: tokens[*curIndex].Literal}
	*curIndex += 1
	return node
}

func parseMathOperator(tokens []*Token, curIndex *int) IExpression {
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
