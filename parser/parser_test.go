package parser

import (
    "testing"
//  "github.com/josketres/monkey-go/ast"
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
    if program == nil {
        t.Fatalf("ParseProgram() returned nil")
    }

}
