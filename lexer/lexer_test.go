package lexer

import (
    "testing"
    "github.com/josketres/monkey-go/token"
)

func TestNextToken(t *testing.T) {
    input := "="

    tests := []struct {
        expectedType token.TokenType
        expectedLiteral string
    } {
        {token.ASSIGN, "="},
        {token.EOF, ""},
    }

    l := New(input)

    for i, tt:= range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
        }
    }
}
