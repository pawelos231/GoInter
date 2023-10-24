package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" //illegal characters
	EOF     = "EOF"     //end of file

	IDENT = "IDENT"
	INT   = "INT"

	//operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Key words
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func lookUpIdentifier(ident string) TokenType {
	if tk, ok := keywords[ident]; ok {
		return tk
	}
	return IDENT
}
