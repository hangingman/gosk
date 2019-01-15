package parser

import (
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/token"
)

// ParseProgram は Parser を受け取ってAST化されたProgramを返す
func (p *Parser) ParseProgram() *ast.Program {
	p.Logger.Debug("ParseProgram!")
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.OPCODE:
		return p.parseMnemonicStatement()
	case token.LABEL:
		return p.parseLabelStatement()
	case token.LBRACKET:
		return p.parseSettingStatement()
	case token.IDENT:
		return p.parseEquStatement()
	default:
		return nil
	}
}

func (p *Parser) parseMnemonicStatement() *ast.MnemonicStatement {

	opcodeParseFn := p.opcodeParseFns[p.curToken.Literal]
	if opcodeParseFn == nil {
		return nil
	}
	stmt := opcodeParseFn()
	p.nextToken()
	p.Logger.Info(stmt.String())
	return stmt
}

func (p *Parser) parseLabelStatement() *ast.LabelStatement {
	stmt := &ast.LabelStatement{
		Token: p.curToken,
		Name:  p.curToken.Literal,
	}
	p.Logger.Info(stmt.String())
	return stmt
}

func (p *Parser) parseSettingStatement() *ast.SettingStatement {

	p.nextToken()

	stmt := &ast.SettingStatement{
		Token: p.curToken,
		Name: &ast.Identifier{
			Token: token.Token{Type: token.SETTING, Literal: string(p.curToken.Literal)},
			Value: p.peekToken.Literal,
		},
	}

	for !p.curTokenIs(token.RBRACKET) && !p.curTokenIs(token.EOF) {
		p.nextToken()
	}

	p.Logger.Info(stmt.String())
	return stmt
}

func (p *Parser) parseEquStatement() *ast.EquStatement {

	if !(p.curTokenIs(token.IDENT) && p.peekTokenIs(token.EQU)) {
		return nil
	}

	stmt := &ast.EquStatement{
		Token: token.Token{Type: token.EQU, Literal: "EQU"},
		Name: &ast.Identifier{
			Token: token.Token{Type: p.curToken.Type, Literal: string(p.curToken.Literal)},
		},
	}

	p.nextToken()
	stmt.Name.Value = p.peekToken.Literal
	p.Logger.Info(stmt.String())
	p.nextToken()
	p.nextToken()

	return stmt
}
