package repl

//temporary file, in final form it will used websocket connection with some frontend

import (
	"bufio"
	"fmt"
	"interpreter/src/lexer"
	"interpreter/src/token"
	"io"
)

// visualization of lexer only, later think about visualizing parsed output in some kind of undersantable way
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
