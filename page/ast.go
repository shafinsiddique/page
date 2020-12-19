package main

type AST struct {
	expressions []IExpression
}

func NewBasicAST () IAST {
	return &AST{}
}

func (ast *AST) GetExpressions()[]IExpression {
	return ast.expressions
}

func (ast *AST) AddExpression(expr IExpression) {
	ast.expressions = append(ast.expressions, expr)
}