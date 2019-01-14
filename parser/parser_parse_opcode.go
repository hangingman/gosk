package parser

import (
	"fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/token"
)

func (p *Parser) parseDBStatement() *ast.MnemonicStatement {

	stmt := &ast.MnemonicStatement{
		Token: p.curToken,
		Name: &ast.Identifier{
			Token: token.Token{Type: token.OPCODE, Literal: string(p.curToken.Literal)},
			Value: p.peekToken.Literal,
		},
	}
	fmt.Printf("parseDBStatement! : %s\n", stmt.String())
	return stmt
}
