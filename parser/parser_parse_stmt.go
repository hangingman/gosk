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
		// stmt := p.parseStatement()
		p.parseStatement()
		// if stmt != nil {
		// 	program.Statements = append(program.Statements, stmt)
		// }
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	fmt.Println("parseStatement")
	switch p.curToken.Type {
	case token.OPCODE:
		return p.parseMnemonicStatement()
	case token.LABEL:
		return p.parseLabelStatement()
	case token.LBRACKET:
		fmt.Println("hei")
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
	fmt.Println("parseSettingStatement")
	stmt := &ast.SettingStatement{Token: p.peekToken}

	for !p.curTokenIs(token.RBRACKET) && !p.curTokenIs(token.EOF) {
		fmt.Println(p.curToken)
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseEquStatement() *ast.EquStatement {
	stmt := &ast.EquStatement{Token: p.curToken}
	return stmt
}
