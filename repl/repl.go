package repl

//temporary file, in final form it will used websocket connection with some frontend

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/token"
	"io"
)

const PROPMT = "->"

func Start(in io.Reader, ouy io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROPMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lexer := lexer.NewLexer(line)
		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
