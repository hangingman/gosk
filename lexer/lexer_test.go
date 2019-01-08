package lexer

import (
	"fmt"
	"testing"
	"github.com/hangingman/gosk/token"
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

	// tests := []struct {
	// 	expectedType    token.TokenType
	// 	expectedLiteral string
	// }{
	// 	{token.EOF, ""},
	// }

	l := New(input)

	// for i, tt := range tests {
	// 	tok := l.NextToken()

	// 	if tok.Type != tt.expectedType {
	// 		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
	// 			i, tt.expectedType, tok.Type)
	// 	}
	// 	if tok.Literal != tt.expectedLiteral {
	// 		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
	// 			i, tt.expectedLiteral, tok.Literal)
	// 	}
	// }

	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		} else {
			fmt.Printf("%s\n", tok)
		}
	}
}
