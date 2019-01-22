package parser

import (
	"fmt"
	"github.com/comail/colog"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/lexer"
	"github.com/hangingman/gosk/token"
	"log"
	"reflect"
)

type (
	opcodeParseFn func() *ast.MnemonicStatement
)

type Parser struct {
	l              *lexer.Lexer
	curIndex       int
	lexedTokens    []token.Token
	errors         []string
	opcodeParseFns map[string]opcodeParseFn // オペコードごとに構文解析を切り替える
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:           l,
		curIndex:    0,
		lexedTokens: make([]token.Token, 50),
		errors:      []string{},
	}
	colog.Register()
	colog.SetDefaultLevel(colog.LInfo)
	colog.SetMinLevel(colog.LInfo)
	colog.SetFlags(log.Lshortfile)

	// オペコードの構文解析方式を格納するmap
	p.opcodeParseFns = make(map[string]opcodeParseFn)
	p.registerOpecode("DB", p.parseDBStatement)
	p.registerOpecode("DW", p.parseDBStatement)
	p.registerOpecode("DD", p.parseDBStatement)
	p.registerOpecode("RESB", p.parseRESBStatement)

	// 初回のTokenを配列に追加
	p.lexedTokens = append(p.lexedTokens, p.curToken())

	// EOFまでトークンを読み込む
	for {
		tok := p.l.NextToken()
		p.lexedTokens = append(p.lexedTokens, tok)
		if tok.Type == token.EOF {
			log.Println(fmt.Sprintf("debug: p.curIndex max = %d", len(p.lexedTokens)))
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

func isNil(x interface{}) bool {
	return x == nil || reflect.ValueOf(x).IsNil()
}

func isNotNil(x interface{}) bool {
	return !isNil(x)
}
