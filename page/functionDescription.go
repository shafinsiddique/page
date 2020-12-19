package main

type FunctionDescription struct {
	name string
	args map[string]int
	tokens []*Token
	inParsing bool
}

func NewFunctionDescription(name string, args []string, tokens[]*Token) *FunctionDescription{
	argsMap := map[string]int{}
	for i,v := range args {
		argsMap[v] = i
	}
	return &FunctionDescription{name: name, args: argsMap, tokens:tokens}
}


