package parser

import (
	"fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/lexer"
	"github.com/hangingman/gosk/token"
	"github.com/sirupsen/logrus"
)

type (
	opcodeParseFn func() *ast.MnemonicStatement
)

type Parser struct {
	l              *lexer.Lexer
	curToken       token.Token
	peekToken      token.Token
	errors         []string
	Logger         *logrus.Logger
	opcodeParseFns map[string]opcodeParseFn // オペコードごとに構文解析を切り替える
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
		Logger: l.Logger,
	}

	// オペコードの構文解析方式を格納するmap
	p.opcodeParseFns = make(map[string]opcodeParseFn)
	p.registerOpecode("DB", p.parseDBStatement)
	p.registerOpecode("DW", p.parseDBStatement)
	p.registerOpecode("DD", p.parseDBStatement)

	// ２つトークンを読み込む。curTokenとpeekTokenの両方がセットされる。
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) registerOpecode(opcode string, fn opcodeParseFn) {
	p.opcodeParseFns[opcode] = fn
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
