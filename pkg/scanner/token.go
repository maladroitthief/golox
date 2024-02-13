package scanner

import "fmt"

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	IDENTIFIER
	STRING
	NUMBER
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF
)

type (
	Token struct {
		tokenType TokenType
		lexeme    string
		literal   string
		line      int
	}

	TokenType int
)

func NewToken(tokenType TokenType, lexeme string, literal string, line int) Token {
	return Token{
		tokenType: tokenType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

func (t Token) ToString() string {
	return fmt.Sprintf("%v %v %v", t.tokenType, t.lexeme, t.literal)
}
