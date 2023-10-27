package lexer

import "interpreter/token"

type Lexer struct {
	input        string
	poistion     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
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

	l.slipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tk = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tk = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case '-':
		tk = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tk = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tk = newToken(token.BANG, l.ch)
		}
	case '/':
		tk = newToken(token.SLASH, l.ch)
	case '*':
		tk = newToken(token.ASTERIK, l.ch)
	case '<':
		tk = newToken(token.LT, l.ch)
	case '>':
		tk = newToken(token.GT, l.ch)
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
	default:
		if isLetter(l.ch) {
			tk.Literal = l.readIdentifier()
			tk.Type = token.LookUpIdentifier(tk.Literal)
			return tk
		} else if isDigit(l.ch) {
			tk.Type = token.INT
			tk.Literal = l.readNumber()
			return tk
		} else {
			tk = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tk
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.poistion
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.poistion]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) slipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.poistion

	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.poistion]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
