package parser

import (
	// "fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/token"
)

// parseDBStatement は DB,DW,DD オペコードを解析する
func (p *Parser) parseDBStatement() *ast.MnemonicStatement {

	stmt := &ast.MnemonicStatement{
		Token: p.curToken,
		Name: &ast.IdentifierArray{
			Token:  token.Token{Type: token.OPCODE, Literal: string(p.curToken.Literal)},
			Values: []string{p.peekToken.Literal},
		},
	}

	for {
		p.nextToken()
		if !(p.peekTokenIs(token.COMMA) || p.peekTokenIs(token.EOF)) {
			break
		}
		p.nextToken()
		stmt.Name.Values = append(stmt.Name.Values, p.peekToken.Literal)
	}

	// fmt.Printf("parseDBStatement! : %s\n", stmt.String())
	return stmt
}
