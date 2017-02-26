package parser

import (
    "github.com/josketres/monkey-go/ast"
    "github.com/josketres/monkey-go/lexer"
    "github.com/josketres/monkey-go/token"
)

type Parser struct {
    l *lexer.Lexer
    curToken token.Token
    peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{l: l}

    // init parser - both curToken and peekToken should be set
    p.nextToken()
    p.nextToken()

    return p
}

func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
    return &ast.Program{}
}

