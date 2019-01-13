package ast

import (
	"github.com/hangingman/gosk/token"
	"strings"
)

// MnemonicStatement は `MOV BX, 15` のような構文を解析する
type MnemonicStatement struct {
	Token token.Token // OPCODE
	Name  *Identifier
	Value Expression
	Line  int
}

func (s *MnemonicStatement) statementNode()       {}
func (s *MnemonicStatement) TokenLiteral() string { return s.Token.Literal }
func (s *MnemonicStatement) String() string       { return "MNEMONIC" }

// SettingStatement は `[FORMAT "WCOFF"]` のような構文を解析する
type SettingStatement struct {
	Token token.Token // SETTING
	Name  *Identifier
	Value string
	Line  int
}

func (s *SettingStatement) statementNode()       {}
func (s *SettingStatement) TokenLiteral() string { return s.Token.Literal }
func (s *SettingStatement) String() string {
	return strings.Join([]string{
		"[",
		token.SETTING + ":",
		s.Name.String(),
		"]"},
		" ")
}

// LabelStatement は `entry:` のような構文を解析する
type LabelStatement struct {
	Token token.Token // LABEL
	Name  *Identifier
	Line  int
}

func (s *LabelStatement) statementNode()       {}
func (s *LabelStatement) TokenLiteral() string { return s.Token.Literal }
func (s *LabelStatement) String() string       { return "LABEL" }

// EquStatement は `BOTPAK  EQU  0x00280000` のような構文を解析する
type EquStatement struct {
	Token token.Token // EQU
	Name  *Identifier
	Value Expression
	Line  int
}

func (s *EquStatement) statementNode()       {}
func (s *EquStatement) TokenLiteral() string { return s.Token.Literal }
func (s *EquStatement) String() string       { return "EQU" }
