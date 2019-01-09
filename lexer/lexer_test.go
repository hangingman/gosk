package lexer

import (
	"github.com/hangingman/gosk/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `[INSTRSET "i486p"]
[BITS 32]
		MOV		EDX,2
		MOV		EBX,msg
		INT		0x40
		MOV		EDX,4
		INT		0x40
msg:
		DB	"hello",0`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// [INSTRSET "i486p"]
		{token.LBRACKET, "["},
		{token.IDENT, "INSTRSET"},
		{token.DOUBLE_QT, "\""},
		{token.IDENT, "i"},
		{token.INT, "486"},
		{token.IDENT, "p"},
		{token.DOUBLE_QT, "\""},
		{token.RBRACKET, "]"},
		// [BITS 32]
		{token.LBRACKET, "["},
		{token.IDENT, "BITS"},
		{token.INT, "32"},
		{token.RBRACKET, "]"},
		// MOV		EDX,2
		{token.IDENT, "MOV"},
		{token.IDENT, "EDX"},
		{token.COMMA, ","},
		{token.INT, "2"},
		// MOV		EBX,msg
		{token.IDENT, "MOV"},
		{token.IDENT, "EBX"},
		{token.COMMA, ","},
		{token.IDENT, "msg"},
		// INT		0x40
		{token.IDENT, "INT"},
		{token.INT, "0"},
		{token.IDENT, "x"},
		{token.INT, "40"},
		// MOV		EDX,4
		{token.IDENT, "MOV"},
		{token.IDENT, "EDX"},
		{token.COMMA, ","},
		{token.INT, "4"},
		// INT		0x40
		{token.IDENT, "INT"},
		{token.INT, "0"},
		{token.IDENT, "x"},
		{token.INT, "40"},
		// msg:
		{token.IDENT, "msg"},
		{token.COLON, ":"},
		// DB	"hello",0
		{token.IDENT, "DB"},
		{token.DOUBLE_QT, "\""},
		{token.IDENT, "hello"},
		{token.DOUBLE_QT, "\""},
		{token.COMMA, ","},
		{token.INT, "0"},
		// EOF!
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
