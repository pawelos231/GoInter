package lexer

import "interpreter/token"

type Lexer struct {
	input        string
	poistion     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}
	lexer.readChar()
	return lexer
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.poistion = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	switch l.ch {
	case '=':
		tk = newToken(token.ASSIGN, l.ch)
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case '(':
		tk = newToken(token.LPAREN, l.ch)
	case ')':
		tk = newToken(token.RPAREN, l.ch)
	case '{':
		tk = newToken(token.LBRACE, l.ch)
	case '}':
		tk = newToken(token.RBRACE, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	case 0:
		tk.Literal = ""
		tk.Type = token.EOF
	}

	l.readChar()
	return tk

}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
