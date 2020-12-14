package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func startRepl(in io.Reader, out io.Writer) {
	fmt.Println("Welcome to Page - A Minimal Lisp Interpreter Written in Go. Please contact @sl2j for questions.")
	prompt := "> "
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, prompt)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		tokens := (&BasicScanner{strScan: line}).GetTokens()
		ast := (&RDParser{}).Parse(tokens)
		ast.Evaluate()

	}
}
func main() {
	startRepl(os.Stdin, os.Stdout)
	//scanner := &BasicScanner{strScan: "(+ 2 (+ 1 1)) (+ 1 1)"}
	//tokens := scanner.GetTokens()
	//parser := &RDParser{}
	//ast := parser.Parse(tokens)
	//ast.Evaluate()
}
