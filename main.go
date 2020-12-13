package main

import "fmt"

//import "fmt"

func main() {
	scanner := &BasicScanner{strScan: " (- 2 3) "}
	tokens := scanner.GetTokens()
	for _, t := range tokens {
		fmt.Println(t.ToString())
	}
	//fmt.Println(scanner.curIndex)
}
