package parser

import (
	"fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/token"
)

// ParseProgram は Parser を受け取ってAST化されたProgramを返す
func (p *Parser) ParseProgram() *ast.Program {
	fmt.Println("ParseProgram!")
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
	fmt.Println("parseStatement!")
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
	stmt := &ast.MnemonicStatement{Token: p.curToken}
	return stmt
}

func (p *Parser) parseLabelStatement() *ast.LabelStatement {
	stmt := &ast.LabelStatement{Token: p.curToken}
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

	fmt.Printf("parseSettingStatement! : %s\n", stmt.String())
	return stmt
}

func (p *Parser) parseEquStatement() *ast.EquStatement {
	stmt := &ast.EquStatement{Token: p.curToken}

	fmt.Printf("parseEquStatement! : %s\n", stmt.String())
	return stmt
}
