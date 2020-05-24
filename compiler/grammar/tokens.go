package token

type TokenType string

type Token struct{
	Type TokenType
	Value string
}

const (
	
	EOF = "EOF"
	ILLEGAL ="ILLEGAL"//illegal token
	SEMICOLON = ";"
	
	LBRACKET ="["//left bracket
	RBRACKET ="]"
	LBRACE = "{"
    RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"
	
    EQ  = "==" //equal sign
	PL_EQ = "+="  //plus equal
    MI_EQ  = "-= "//minus equal
	PP = "++"  //plus plus
    MM = "--" //minus minus
    NE = "!=" //not equal
	MULT_EQ = "*="
	DEVIDE_EQ = "/="
	POWER_EQ = "*="

    AND = "&& "// and &&
    OR = "||"  //or ||
	
    IF = "if"//if
    ELSE = "else"//else
   
	FOR  = "for"//for
    WHILE = "while"//while
    FOREACH  = "foreach"//for each
	
	TRUE = "true"
	FALSE = "false"


	FUNCTION = "FUNCTION"//function

	ASIGN = "="
	PLUS = "+"
	MINUS = "-"
	MULTIPLY = "*"
	DIVISION = "/"
	NOT = "!"
	POWER = "^"

	IDENT = "IDENT"
	INT = "INT"
	STRING = "STRING"
	RETURN = "RETURN"
	
	COMMA = ","
	DOT = "."

	SMALLER = "<"
	BIGGER = ">"
	SMALLEREQUAL = "<="
	BIGGEREQUAL = ">="
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
