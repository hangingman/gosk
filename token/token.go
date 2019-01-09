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
	LT        = "<"
	GT        = ">"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	DOUBLE_QT = "\""
	EQU       = "EQU"
	GLOBAL    = "GLOBAL"
)

var keywords = map[string]TokenType{
	"EQU":    EQU,
	"GLOBAL": GLOBAL,
}

func LookupIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
