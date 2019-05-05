package ast

import (
	"github.com/hangingman/gosk/token"
)

// MnemonicStatement は `MOV BX, 15` のような構文を解析する
type MnemonicStatement struct {
	Token    token.Token // OPCODE
	Name     *IdentifierArray
	NextNode *Statement
	PrevNode *Statement
}

func (s *MnemonicStatement) statementNode()       {}
func (s *MnemonicStatement) TokenLiteral() string { return s.Token.Literal }
func (s *MnemonicStatement) String() string {
	return "{ " + token.OPCODE + ":" + s.Name.String() + " }"
}
func (s *MnemonicStatement) nextNode() *Statement { return s.NextNode }
func (s *MnemonicStatement) prevNode() *Statement { return s.PrevNode }

// SettingStatement は `[FORMAT "WCOFF"]` のような構文を解析する
type SettingStatement struct {
	Token    token.Token // SETTING
	Name     *Identifier
	Value    string
	NextNode *Statement
	PrevNode *Statement
}

func (s *SettingStatement) statementNode()       {}
func (s *SettingStatement) TokenLiteral() string { return s.Token.Literal }
func (s *SettingStatement) String() string {
	return "{ " + token.SETTING + ":" + s.Name.String() + " }"
}
func (s *SettingStatement) nextNode() *Statement { return s.NextNode }
func (s *SettingStatement) prevNode() *Statement { return s.PrevNode }

// LabelStatement は `entry:` のような構文を解析する
type LabelStatement struct {
	Token    token.Token // LABEL
	Name     string
	NextNode *Statement
	PrevNode *Statement
}

func (s *LabelStatement) statementNode()       {}
func (s *LabelStatement) TokenLiteral() string { return s.Token.Literal }
func (s *LabelStatement) String() string {
	return "{ " + token.LABEL + ": \"" + s.Name + "\" }"
}
func (s *LabelStatement) nextNode() *Statement { return s.NextNode }
func (s *LabelStatement) prevNode() *Statement { return s.PrevNode }

// EquStatement は `BOTPAK  EQU  0x00280000` のような構文を解析する
type EquStatement struct {
	Token    token.Token // EQU
	Name     *Identifier
	Value    string // TODO: 後でExpressionにするかも
	NextNode *Statement
	PrevNode *Statement
}

func (s *EquStatement) statementNode()       {}
func (s *EquStatement) TokenLiteral() string { return s.Token.Literal }
func (s *EquStatement) String() string {
	return "{ " + token.EQU + ":" + s.Name.String() + " }"
}
func (s *EquStatement) nextNode() *Statement { return s.NextNode }
func (s *EquStatement) prevNode() *Statement { return s.PrevNode }

type DummyStatement struct{}

func (ds *DummyStatement) String() string       { return "dummy" }
func (ds *DummyStatement) TokenLiteral() string { return "dummy-token-literal" }
func (ds *DummyStatement) statementNode()       {}
