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
	curIndex       int
	lexedTokens    []token.Token
	errors         []string
	Logger         *logrus.Logger
	opcodeParseFns map[string]opcodeParseFn // オペコードごとに構文解析を切り替える
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:           l,
		curIndex:    0,
		lexedTokens: make([]token.Token, 50),
		errors:      []string{},
		Logger:      l.Logger,
	}

	// オペコードの構文解析方式を格納するmap
	p.opcodeParseFns = make(map[string]opcodeParseFn)
	p.registerOpecode("DB", p.parseDBStatement)
	p.registerOpecode("DW", p.parseDBStatement)
	p.registerOpecode("DD", p.parseDBStatement)

	// 初回のTokenを配列に追加
	p.lexedTokens = append(p.lexedTokens, p.curToken())
	// p.lexedTokens = []token.Token{p.curToken()}

	// EOFまでトークンを読み込む
	for {
		tok := p.l.NextToken()
		p.lexedTokens = append(p.lexedTokens, tok)
		if tok.Type == token.EOF {
			p.Logger.Info(fmt.Sprintf("p.curIndex max = %d", len(p.lexedTokens)))
			break
		}
	}
	return p
}

func (p *Parser) registerOpecode(opcode string, fn opcodeParseFn) {
	p.opcodeParseFns[opcode] = fn
}

func (p *Parser) maxTokenSize() int {
	return len(p.lexedTokens)
}

func (p *Parser) curToken() token.Token {
	if p.curIndex >= p.maxTokenSize() {
		return token.Token{Type: token.EOF, Literal: ""}
	}
	return p.lexedTokens[p.curIndex]
}

func (p *Parser) peekToken() token.Token {
	if p.curIndex+1 >= p.maxTokenSize() {
		return token.Token{Type: token.EOF, Literal: ""}
	}
	return p.lexedTokens[p.curIndex+1]
}

func (p *Parser) lookAhead(n int) token.Token {
	if p.curIndex+n < p.maxTokenSize() {
		return p.lexedTokens[p.curIndex+n]
	}
	return token.Token{Type: token.ILLEGAL, Literal: ""}
}

func (p *Parser) nextToken() {
	p.curIndex++
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken().Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken().Type == t
}

func (p *Parser) lookAheadIs(n int, t token.TokenType) bool {
	return p.lookAhead(n).Type == t
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
		t, p.peekToken().Type)
	p.errors = append(p.errors, msg)
}
