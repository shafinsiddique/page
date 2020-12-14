package main

type AST struct {
	expressions []IExpression
}

func NewBasicAST () IAST {
	return &AST{}
}

func (ast *AST) ToString() string {
	str := ""
	for _, expr := range ast.expressions {
		str += expr.ToString() + "\n"
	}

	return str
}


func (ast *AST) GetExpressions()[]IExpression {
	return ast.expressions
}

func (ast *AST) AddExpression(expr IExpression) {
	ast.expressions = append(ast.expressions, expr)
}