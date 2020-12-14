package main

import (
	"fmt"
	"io"
	"os"
)

func startRepl(in io.Reader, out io.Writer) {
	fmt.Println("Welcome to Page - A Minimal Lisp Interpreter Written in Go. \nFind more information on our " +
		"github repository.")

	//prompt := "> "
	//scanner := bufio.NewScanner(in)
	tokenizer := NewBasicScanner()
	tokens := tokenizer.GetTokens("(+ 1 4) (+ 1 2)")
	for _, t := range tokens {
		fmt.Println(t.ToString())
	}
	//for {
	//	fmt.Fprint(out, prompt)
	//	scanned := scanner.Scan()
	//	if !scanned {
	//		return
	//	}
	//
	//	line := scanner.Text()
	//
	//}
}
func main() {
	startRepl(os.Stdin, os.Stdout)
}
