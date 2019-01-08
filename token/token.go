package token

// TokenType は stringのエイリアス
type TokenType string

// Token は文字列をトークン化した後の情報を保持する
// TokenType にはtoken内で定義したconst値
// Literal には実際に読み取った文字列
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	SHARP     = "#"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NOTEQ     = "!="
)

var keywords = map[string]TokenType{
	"EQU": EQU,
}

func LookupIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
