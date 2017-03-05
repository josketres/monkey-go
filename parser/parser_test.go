package parser

import (
    "testing"
    "github.com/josketres/monkey-go/ast"
    "github.com/josketres/monkey-go/lexer"
)

func TestLetStatements(t *testing.T) {
    input := `
    let x = 5;
    let y = 10;
    let foobar = 838383;
    `

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParserErrors(t, p)
    if program == nil {
        t.Fatalf("ParseProgram() returned nil")
    }
    if length := len(program.Statements); length != 3 {
        t.Fatalf("program.Statements does not contain 3 statements. got=%d", length)
    }

    tests := []struct {
        expectedIdentifier string
    }{
        {"x"},
        {"y"},
        {"foobar"},
    }

    for i, tt := range tests {
        stmt := program.Statements[i]
        if !testLetStatement(t, stmt, tt.expectedIdentifier) {
            return
        }
    }
}

func checkParserErrors(t *testing.T, p *Parser) {
    errors := p.Errors()
    if len(errors) == 0 {
        return
    }

    t.Errorf("parser has %d errors", len(errors))
    for _, msg := range errors {
        t.Errorf("parser error: %q", msg)
    }
    t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, expectedIdentifier string) bool {
    if tl := s.TokenLiteral(); tl != "let" {
        t.Errorf("s.TokenLiteral() not 'let'. got=%q", tl)
        return false
    }

    letStmt, ok := s.(*ast.LetStatement)
    if !ok {
        t.Errorf("s not *ast.LetStatement. got=%T", s)
        return false
    }

    if n := letStmt.Name.Value; n != expectedIdentifier {
        t.Errorf("letStmt.Name.Value not '%s'. got=%s", expectedIdentifier, n)
        return false
    }

    if tl := letStmt.Name.TokenLiteral(); tl != expectedIdentifier {
        t.Errorf("s.Name not '%s'. got=%s", expectedIdentifier, tl)
        return false
    }

    return true
}
