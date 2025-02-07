package token

type TokenType string

type Token struct{
   Type TokenType
   Literal string
}

const(
   ILLEGAL = "ILLEGAL"
   EOF = "EOF"

   //identifier + literals
   IDENT = "IDENT"
   INT = "INT"

   //operators
   ASSIGN = "="
   PLUS = "+"
   MINUS = "-"
   BANG = "!"
   ASTERISK = "*"
   SLASH = "/"
   //delimiter
   COMMA = ","
   SEMICOLON = ";"

   LPAREN = "("
   RPAREN = ")"
   LBRACE = "{"
   RBRACE = "}"

   LT = "<"
   GT = ">"

   EQ = "=="
   NOT_EQ = "!="
   //keywords
   FUNCTION = "FUNCTION"
   LET = "LET"

   IF = "IF"
   TRUE = "TRUE"
   FALSE = "FALSE"
   ELSE = "ELSE"
   RETURN = "RETURN"
)

var keywords = map[string]TokenType{
   "fn": FUNCTION,
   "let": LET,
   "if" : IF,
   "true": TRUE,
   "false": FALSE,
   "else": ELSE,
   "return": RETURN,
}

func LookupIdent(ident string) TokenType{
   if tok, ok := keywords[ident]; ok{
      return tok
   }
   return IDENT
}
