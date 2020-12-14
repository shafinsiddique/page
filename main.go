package main

func main() {
	scanner := &BasicScanner{strScan: "(+2 (+ 1 1))"}
	tokens := scanner.GetTokens()
	parser := &RDParser{}
	ast := parser.Parse(tokens)
	ast.Evaluate()
	//fmt.Println(ast.ToString())
	//parser.Parse(tokens)


}
//	for _, t := range tokens {
//		fmt.Println(t.ToString())
//	}
//}
