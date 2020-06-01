package token

type TokenType string

type Token struct{
	Type TokenType
	Value string
}

const (
	
	EOF = "EOF"
	ILLEGAL ="ILLEGAL"

	LBRACKET ="["
	RBRACKET ="]"
	LBRACE = "{"
    RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"
	
    EQ  = "==" 
	PL_EQ = "+="  
    MI_EQ  = "-= "
	PP = "++" 
    MM = "--"
    NE = "!="
	MULT_EQ = "*="
	DEVIDE_EQ = "/="
	POW_EQ = "^="
	
    AND = "&& "
    OR = "||" 
	
    IF = "if"
    ELSE = "else"
   
	FOR  = "for"
    WHILE = "while"
    FOREACH  = "foreach"
	
	TRUE = "true"
	FALSE = "false"

	FUNCTION = "FUNCTION"
	LET = "LET"

	ASIGN = "="
	PLUS = "+"
	MINUS = "-"
	MULTIPLY = "*"
	DIVISION = "/"
	NOT = "!"
	POW = "^"
	SMALLER = "<"
	BIGGER = ">"
	SMALL_EQ = "<="
	BIG_EQ = ">="

	IDENT = "IDENT"
	INT = "INT"
	STRING = "STRING"
	RETURN = "RETURN"
	
	COMMA = ","
	DOT = "."
	SEMICOLON = ";"

)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let":LET,
	"if":IF,
	"else":ELSE,
	"while":WHILE,
	"for":FOR,
	"foreach":FOREACH,
	"true":TRUE,
	"false":FALSE,
	"return":RETURN}

func LookUpIdent(ident string) TokenType{
	if tok, ok := keywords[ident];ok{
		return tok
	}
	return IDENT
}
