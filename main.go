package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func startRepl(in io.Reader, out io.Writer) {
	fmt.Println("Welcome to Page - A Minimal Lisp Interpreter Written in Go. \nFind more information on our " +
		"github repository.")

	tokenizer := NewBasicScanner()
	ast := NewBasicAST()
	parser := NewRDParser()
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
		parser.Parse(ast, tokens)
		evaluator.EvaluateLast(ast)
	}
}

func main() {
	log.SetFlags(0)
	startRepl(os.Stdin, os.Stdout)
}
