package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers + literals
	IDENT = "IDENT" //add, foobar, x ,y, etc
	INT   = "INT"   //12345

	//operators
	ASSIGN = "="
	PLUS   = "+"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}
