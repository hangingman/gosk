package parser

import (
	"fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/token"
	"log"
)

// parseDBStatement は DB,DW,DD オペコードを解析する
func (p *Parser) parseDBStatement() *ast.MnemonicStatement {
	log.Println(fmt.Sprintf("debug: Parser: cur = %s, peek = %s, peek+1 = %s",
		p.curToken(),
		p.peekToken(),
		p.lookAhead(2),
	))

	stmt := &ast.MnemonicStatement{
		Token: p.curToken(),
		Name: &ast.IdentifierArray{
			Tokens: []token.Token{p.curToken()},
			Values: []string{p.curToken().Literal},
		},
	}
	p.nextToken()

	//    [-]    [0]  [1]
	// <OPCODE> <IMM> <,> ...
	if p.peekTokenIs(token.COMMA) {
		//  [-]      [0]  [1]  [2]  [3]
		// <OPCODE> <IMM> <,> <IMM> <,> ...
		for p.peekTokenIs(token.COMMA) {
			stmt.Name.Tokens = append(stmt.Name.Tokens, p.curToken())
			stmt.Name.Values = append(stmt.Name.Values, p.curToken().Literal)
			p.nextToken()
			p.nextToken()
		}
		if p.curTokenIs(token.INT) || p.curTokenIs(token.HEX_LIT) || p.curTokenIs(token.STR_LIT) {
			stmt.Name.Tokens = append(stmt.Name.Tokens, p.curToken())
			stmt.Name.Values = append(stmt.Name.Values, p.curToken().Literal)
		}

	} else {
		//    [-]    [0]
		// <OPCODE> <IMM>
		stmt.Name.Tokens = append(stmt.Name.Tokens, p.curToken())
		stmt.Name.Values = append(stmt.Name.Values, p.curToken().Literal)
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

	if p.lookAheadIs(2, token.MINUS) && p.lookAheadIs(3, token.DOLLAR) {
		p.nextToken()
		p.nextToken()
		stmt.Name.Tokens = append(stmt.Name.Tokens, p.curToken())
		stmt.Name.Values = append(stmt.Name.Values, p.curToken().Literal)
		p.nextToken()
		stmt.Name.Tokens = append(stmt.Name.Tokens, p.curToken())
		stmt.Name.Values = append(stmt.Name.Values, p.curToken().Literal)
	}

	return stmt
}

// parseORGStatement は ORG オペコードを解析する
func (p *Parser) parseORGStatement() *ast.MnemonicStatement {

	stmt := &ast.MnemonicStatement{
		Token: p.curToken(),
		Name: &ast.IdentifierArray{
			Tokens: []token.Token{p.curToken()},
			Values: []string{p.curToken().Literal},
		},
	}
	// ORG 命令の後は必ずintリテラルかhexリテラル
	if !p.expectPeek(token.INT) && !p.expectPeek(token.HEX_LIT) {
		return nil
	}
	stmt.Name.Tokens = append(stmt.Name.Tokens, p.curToken())
	stmt.Name.Values = append(stmt.Name.Values, p.curToken().Literal)

	return stmt
}

// parseOnlyOpcodeStatement は オペコードのみの文を解析する
func (p *Parser) parseOnlyOpcodeStatement() *ast.MnemonicStatement {

	stmt := &ast.MnemonicStatement{
		Token: p.curToken(),
		Name: &ast.IdentifierArray{
			Tokens: []token.Token{p.curToken()},
			Values: []string{p.curToken().Literal},
		},
	}
	return stmt
}

// parseJMPStatement は JMP オペコードを解析する
func (p *Parser) parseJMPStatement() *ast.MnemonicStatement {

	stmt := &ast.MnemonicStatement{
		Token: p.curToken(),
		Name: &ast.IdentifierArray{
			Tokens: []token.Token{p.curToken()},
			Values: []string{p.curToken().Literal},
		},
	}
	// JMP 命令の後は必ず識別子
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name.Tokens = append(stmt.Name.Tokens, p.curToken())
	stmt.Name.Values = append(stmt.Name.Values, p.curToken().Literal)

	return stmt
}

// parseMOVStatement は MOV オペコードを解析する
func (p *Parser) parseMOVStatement() *ast.MnemonicStatement {
	// MOV DST  ,  SRC
	// [0] [1] [2] [3]
	stmt := &ast.MnemonicStatement{
		Token: p.curToken(),
		Name: &ast.IdentifierArray{
			Tokens: []token.Token{p.curToken()},
			Values: []string{p.curToken().Literal},
		},
	}

	p.nextToken()
	// MOV DST  ,  SRC
	// [-] [0] [1] [2]
	stmt.Name.Tokens = append(stmt.Name.Tokens, p.curToken())
	stmt.Name.Values = append(stmt.Name.Values, p.curToken().Literal)

	for {
		// 最初のカンマまで読み取る
		if p.curTokenIs(token.COMMA) {
			break
		}
		if p.curTokenIs(token.EOF) {
			return nil
		}
		p.nextToken()
		stmt.Name.Tokens = append(stmt.Name.Tokens, p.peekToken())
		stmt.Name.Values = append(stmt.Name.Values, p.peekToken().Literal)
	}

	if p.peekTokenIs(token.LBRACKET) {
		// もしブラケットを見つけたら右側の終端まで読み取る
		for {
			if p.peekTokenIs(token.RBRACKET) {
				break
			}
			if p.curTokenIs(token.EOF) {
				return nil
			}
			p.nextToken()
			stmt.Name.Tokens = append(stmt.Name.Tokens, p.peekToken())
			stmt.Name.Values = append(stmt.Name.Values, p.peekToken().Literal)
		}
	}

	return stmt
}
