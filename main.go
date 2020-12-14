package main

import "fmt"

func main() {
	scanner := &BasicScanner{strScan: " (- 2 3) "}
	tokens := scanner.GetTokens()
	parser := &RDParser{}
	ast := parser.Parse(tokens)
	fmt.Println(ast.ToString())
	//parser.Parse(tokens)


}
//	for _, t := range tokens {
//		fmt.Println(t.ToString())
//	}
//}
