package main

type RecursiveCallNode struct {
	tokens []*Token
	fds map[string]*FunctionDescription
}

func (node RecursiveCallNode) Evaluate()interface{} {
	parser := RDParser{}
	ast := AST{}
	parser.Parse(&ast, node.tokens, node.fds)
	expressions := ast.GetExpressions()
	return expressions[len(expressions)-1].Evaluate()
}

func (node RecursiveCallNode) GetType()ExpressionType {
	return BOOLEAN_EXPR
}
