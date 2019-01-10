package ast

import (
	"bytes"
	"github.com/hangingman/gosk/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.Token
	Value string
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// MnemonicStatement は `MOV BX, 15` のような構文を解析する
type MnemonicStatement struct {
	Token token.Token // OPCODE
	Name  *Identifier
	Value Expression
	Line  int
}

// SettingStatement は `[FORMAT "WCOFF"]` のような構文を解析する
type SettingStatement struct {
	Token token.Token // SETTING
	Name  *Identifier
	Value Expression
	Line  int
}

// LabelStatement は `entry:` のような構文を解析する
type LabelStatement struct {
	Token token.Token // LABEL
	Name  *Identifier
	Line  int
}

// EquStatement は `BOTPAK  EQU  0x00280000` のような構文を解析する
type EquStatement struct {
	Token token.Token // EQU
	Name  *Identifier
	Value Expression
	Line  int
}
