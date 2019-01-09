package lexer

import (
	"fmt"
	"github.com/hangingman/gosk/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

type LexerTest struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func testLexerParsedTokens(t *testing.T, tok token.Token, tt *LexerTest, idx int) {

	assert.Equal(t, tok.Type, tt.expectedType,
		fmt.Sprintf("tests[%d] - tokentype wrong. expected=%q, got=%q", idx, tt.expectedType, tok.Type))

	assert.Equal(t, tok.Literal, tt.expectedLiteral,
		fmt.Sprintf("tests[%d] - literal wrong. expected=%q, got=%q", idx, tt.expectedLiteral, tok.Literal))
}

func TestCommentLines(t *testing.T) {
	input := `VBEMODE	EQU		0x105			; 1024 x  768 x 8bitカラー`

	tests := []LexerTest{
		{token.IDENT, "VBEMODE"},
		{token.EQU, "EQU"},
		{token.INT, "0"},
		{token.IDENT, "x"},
		{token.INT, "105"},

		// EOF!
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		testLexerParsedTokens(t, tok, &tt, i)
	}
}

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

	tests := []LexerTest{
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
		testLexerParsedTokens(t, tok, &tt, i)
	}
}
