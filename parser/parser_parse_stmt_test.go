package parser

import (
	"fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/lexer"
	"github.com/hangingman/gosk/token"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"testing"
)

func TestAsmHead(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	asmheadPath := path.Join(path.Dir(filename), "..", "testdata", "asmhead.nas")
	err := os.Chdir("../../")

	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadFile(asmheadPath)
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	// fmt.Println(input)
	l := lexer.New(input)

	for {
		tok := l.NextToken()
		// fmt.Printf("%s\n", tok)
		if tok.Type == token.EOF {
			break
		}
	}

	p := New(l)
	program := p.ParseProgram()
	for stmt := range program.Statements {
		fmt.Printf("%T\n", stmt)
	}
}

func TestParseSettingStatement(t *testing.T) {
	input := `[INSTRSET "i486p"]`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	// 取得できる Statement は１つ
	assert.Equal(t, len(program.Statements), 1)
	stmt, ok := program.Statements[0].(*ast.SettingStatement)
	// Statement は SettingStatement
	assert.True(t, ok)
	assert.NotNil(t, stmt)
	// トークンの中身
	assert.Equal(t,
		token.Token{Type: token.SETTING, Literal: "INSTRSET"},
		stmt.Token)

	ident := stmt.Name
	assert.Equal(t,
		token.Token{Type: token.SETTING, Literal: "INSTRSET"},
		ident.Token)
	assert.Equal(t,
		"\"i486p\"",
		ident.Value)
}

func TestParseEquStatement(t *testing.T) {
	input := `BOTPAK	EQU		0x00280000		; bootpackのロード先`
	l := lexer.New(input)
	p := New(l)
	// fmt.Println(p)
	program := p.ParseProgram()
	// 取得できる Statement は１つ
	assert.Equal(t, len(program.Statements), 1)
	stmt, ok := program.Statements[0].(*ast.EquStatement)
	// Statement は EquStatement
	assert.True(t, ok)
	assert.NotNil(t, stmt)
	// トークンの中身
	assert.Equal(t,
		token.Token{Type: token.EQU, Literal: "EQU"},
		stmt.Token)
	ident := stmt.Name
	assert.Equal(t,
		token.Token{Type: token.IDENT, Literal: "BOTPAK"},
		ident.Token)
	assert.Equal(t,
		"0x00280000",
		ident.Value)
}
