package lexer

import (
	"interpreter-in-go/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

    expectedTokens := []struct {
		Type    token.TokenType
		Literal string
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

	lexer := New(input)

	for i, expectedToken := range expectedTokens {
		token := lexer.NextToken()

		if token.Type != expectedToken.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, expectedToken.Type, token.Type)
		}
		if token.Literal != expectedToken.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, expectedToken.Literal, token.Literal)
		}
	}
}
