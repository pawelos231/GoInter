package token

type TokenType string

// Literal = let, if, fn, or for IDENT, it can be any variable name like: myVar, anotherVar
// TokenType = IDENT, FUNCTION, LET...
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
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	BANG    = "!"
	SLASH   = "/"
	ASTERIK = "*"
	LT      = "<"
	GT      = ">"
	EQ      = "=="
	NOT_EQ  = "!="

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookUpIdentifier(ident string) TokenType {
	if tk, ok := keywords[ident]; ok {
		return tk
	}
	return IDENT
}
