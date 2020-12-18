package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func startRepl(in io.Reader, out io.Writer) {
	fmt.Println("\nWelcome to Page: A Functional Programming Language based on Lisp.\nVerson : 1.0\n" +
		"Written In: Go\nBuilt By: Shafin\n")

	tokenizer := NewBasicScanner()
	ast := NewBasicAST()
	parser := NewRDParser()
	fds := map[string]*FunctionDescription{}
	evaluator := NewEvaluator()
	prompt := "> "
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		tokens := tokenizer.GetTokens(line)
		parser.Parse(ast, tokens, fds)
		evaluator.EvaluateLast(ast)
	}
}

func main() {
	log.SetFlags(0)
	startRepl(os.Stdin, os.Stdout)
}
