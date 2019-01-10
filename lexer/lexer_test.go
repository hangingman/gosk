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
	assert.Equal(t, tt.expectedType, tok.Type,
		fmt.Sprintf("tests[%d] - tokentype wrong. expected=%q, got=%q", idx, tt.expectedType, tok.Type))

	assert.Equal(t, tt.expectedLiteral, tok.Literal,
		fmt.Sprintf("tests[%d] - literal wrong. expected=%q, got=%q", idx, tt.expectedLiteral, tok.Literal))
}

func TestCommentLines(t *testing.T) {
	input := `VBEMODE	EQU		0x105			; 1024 x  768 x 8bitカラー`

	tests := []LexerTest{
		{token.IDENT, "VBEMODE"},
		{token.EQU, "EQU"},
		{token.HEX_LIT, "0x105"},

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
		{token.SETTING, "INSTRSET"},
		{token.STR_LIT, "\"i486p\""},
		{token.RBRACKET, "]"},
		// [BITS 32]
		{token.LBRACKET, "["},
		{token.SETTING, "BITS"},
		{token.INT, "32"},
		{token.RBRACKET, "]"},
		// MOV		EDX,2
		{token.OPCODE, "MOV"},
		{token.IDENT, "EDX"},
		{token.COMMA, ","},
		{token.INT, "2"},
		// MOV		EBX,msg
		{token.OPCODE, "MOV"},
		{token.IDENT, "EBX"},
		{token.COMMA, ","},
		{token.IDENT, "msg"},
		// INT		0x40
		{token.OPCODE, "INT"},
		{token.HEX_LIT, "0x40"},
		// MOV		EDX,4
		{token.OPCODE, "MOV"},
		{token.IDENT, "EDX"},
		{token.COMMA, ","},
		{token.INT, "4"},
		// INT		0x40
		{token.OPCODE, "INT"},
		{token.HEX_LIT, "0x40"},
		// msg:
		{token.LABEL, "msg:"},
		// DB	"hello",0
		{token.OPCODE, "DB"},
		{token.STR_LIT, "\"hello\""},
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
