package ast

import (
	"github.com/hangingman/gosk/object"
	"github.com/hangingman/gosk/token"
)

// MnemonicStatement は `MOV BX, 15` のような構文を解析する
type MnemonicStatement struct {
	Token    token.Token // OPCODE
	Name     *IdentifierArray
	NextNode Statement
	PrevNode Statement
	Bin      *object.Binary
}

func (s *MnemonicStatement) statementNode()       {}
func (s *MnemonicStatement) TokenLiteral() string { return s.Token.Literal }
func (s *MnemonicStatement) String() string {
	name := "nil"
	if isNotNil(s) && isNotNil(s.Name) {
		name = s.Name.String()
	}
	return "{ " + token.OPCODE + ":" + name + " }"
}
func (s *MnemonicStatement) SetNextNode(stmt Statement) { s.NextNode = stmt }
func (s *MnemonicStatement) SetPrevNode(stmt Statement) { s.PrevNode = stmt }
func (s *MnemonicStatement) GetNextNode() Statement     { return s.NextNode }
func (s *MnemonicStatement) GetPrevNode() Statement     { return s.PrevNode }

// SettingStatement は `[FORMAT "WCOFF"]` のような構文を解析する
type SettingStatement struct {
	Token    token.Token // SETTING
	Name     *Identifier
	Value    string
	NextNode Statement
	PrevNode Statement
}

func (s *SettingStatement) statementNode()       {}
func (s *SettingStatement) TokenLiteral() string { return s.Token.Literal }
func (s *SettingStatement) String() string {
	return "{ " + token.SETTING + ":" + s.Name.String() + " }"
}
func (s *SettingStatement) SetNextNode(stmt Statement) { s.NextNode = stmt }
func (s *SettingStatement) SetPrevNode(stmt Statement) { s.PrevNode = stmt }
func (s *SettingStatement) GetNextNode() Statement     { return s.NextNode }
func (s *SettingStatement) GetPrevNode() Statement     { return s.PrevNode }

// LabelStatement は `entry:` のような構文を解析する
type LabelStatement struct {
	Token    token.Token // LABEL
	Name     string
	NextNode Statement
	PrevNode Statement
}

func (s *LabelStatement) statementNode()       {}
func (s *LabelStatement) TokenLiteral() string { return s.Token.Literal }
func (s *LabelStatement) String() string {
	return "{ " + token.LABEL + ": \"" + s.Name + "\" }"
}
func (s *LabelStatement) SetNextNode(stmt Statement) { s.NextNode = stmt }
func (s *LabelStatement) SetPrevNode(stmt Statement) { s.PrevNode = stmt }
func (s *LabelStatement) GetNextNode() Statement     { return s.NextNode }
func (s *LabelStatement) GetPrevNode() Statement     { return s.PrevNode }

// EquStatement は `BOTPAK  EQU  0x00280000` のような構文を解析する
type EquStatement struct {
	Token    token.Token // EQU
	Name     *Identifier
	Value    token.Token
	NextNode Statement
	PrevNode Statement
}

func (s *EquStatement) statementNode()       {}
func (s *EquStatement) TokenLiteral() string { return s.Token.Literal }
func (s *EquStatement) String() string {
	return "{ " + token.EQU + ":" + s.Name.String() + " }"
}
func (s *EquStatement) SetNextNode(stmt Statement) { s.NextNode = stmt }
func (s *EquStatement) SetPrevNode(stmt Statement) { s.PrevNode = stmt }
func (s *EquStatement) GetNextNode() Statement     { return s.NextNode }
func (s *EquStatement) GetPrevNode() Statement     { return s.PrevNode }

type DummyStatement struct{}

func (ds *DummyStatement) String() string       { return "dummy" }
func (ds *DummyStatement) TokenLiteral() string { return "dummy-token-literal" }
func (ds *DummyStatement) statementNode()       {}
