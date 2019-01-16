package parser

import (
	// "fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/token"
)

// parseDBStatement は DB,DW,DD オペコードを解析する
func (p *Parser) parseDBStatement() *ast.MnemonicStatement {

	stmt := &ast.MnemonicStatement{
		Token: p.curToken(),
		Name: &ast.IdentifierArray{
			Tokens: []token.Token{p.curToken()},
			Values: []string{p.curToken().Literal},
		},
	}

	for {
		stmt.Name.Tokens = append(stmt.Name.Tokens, p.peekToken())
		stmt.Name.Values = append(stmt.Name.Values, p.peekToken().Literal)
		if !(p.peekTokenIs(token.COMMA) || p.peekTokenIs(token.EOF)) {
			break
		}
		p.nextToken()
		p.nextToken()
	}
	return stmt
}
