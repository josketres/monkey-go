package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILEGAL"
    EOF     = "EOF"

    // Identifiers + literals
    IDENT = "IDENT" // add, foobar, x, y, ...
    INT   = "INT"   // 1234566

    // Operators
    ASSIGN   = "="
    PLUS     = "+"
    BANG     = "!"
    MINUS    = "-"
    SLASH    = "/"
    ASTERISK = "*"
    LT       = "<"
    GT       = ">"

    // Delimiters
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET      = "LET"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"
    TRUE     = "TRUE"
    FALSE    = "FALSE"
)

var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
    "true": TRUE,
    "false": FALSE,
}

func LookupIdent(ident string) TokenType {
    if tokenType, ok := keywords[ident]; ok {
        return tokenType
    }
    return IDENT
}
