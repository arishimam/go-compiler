package lexer

import (
	"../token"
	"testing"
)

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position
	ch           byte // current char under position
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func TestNextToken(t *testing.T) { // lookup this parameter
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests { // lookup the syntax of this for loop
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
