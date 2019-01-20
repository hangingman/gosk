package parser

import (
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/token"
)

// parseDBStatement は DB,DW,DD オペコードを解析する
func (p *Parser) parseDBStatement() *ast.MnemonicStatement {
	p.logger.DebugF("Parser: cur = %s, peek = %s", p.curToken(), p.peekToken())

	stmt := &ast.MnemonicStatement{
		Token: p.curToken(),
		Name: &ast.IdentifierArray{
			Tokens: []token.Token{p.curToken()},
			Values: []string{p.curToken().Literal},
		},
	}

	//    [0]    [1]  [2]
	// <OPCODE> <IMM> <,> ...
	if p.lookAheadIs(2, token.COMMA) {
		//  [0]	 [1]  [2]  [3]
		// <IMM> <,> <IMM> <,> ...
		for p.lookAheadIs(1, token.COMMA) {
			stmt.Name.Tokens = append(stmt.Name.Tokens, p.curToken())
			stmt.Name.Values = append(stmt.Name.Values, p.curToken().Literal)
			p.nextToken()
			p.nextToken()
		}
		stmt.Name.Tokens = append(stmt.Name.Tokens, p.curToken())
		stmt.Name.Values = append(stmt.Name.Values, p.curToken().Literal)
	} else {
		//    [0]    [1]
		// <OPCODE> <IMM>
		stmt.Name.Tokens = append(stmt.Name.Tokens, p.peekToken())
		stmt.Name.Values = append(stmt.Name.Values, p.peekToken().Literal)
	}

	return stmt
}

// parseRESBStatement は RESB オペコードを解析する
func (p *Parser) parseRESBStatement() *ast.MnemonicStatement {

	stmt := &ast.MnemonicStatement{
		Token: p.curToken(),
		Name: &ast.IdentifierArray{
			Tokens: []token.Token{p.curToken()},
			Values: []string{p.curToken().Literal},
		},
	}
	stmt.Name.Tokens = append(stmt.Name.Tokens, p.peekToken())
	stmt.Name.Values = append(stmt.Name.Values, p.peekToken().Literal)

	return stmt
}
