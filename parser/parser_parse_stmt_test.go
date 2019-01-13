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
	assert.Equal(t, stmt.Token,
		token.Token{Type: token.SETTING, Literal: "INSTRSET"})
	ident := stmt.Name
	assert.Equal(t, ident.Token,
		token.Token{Type: token.SETTING, Literal: "INSTRSET"})
	assert.Equal(t, ident.Value, "\"i486p\"")
}

func TestParseEquStatement(t *testing.T) {
	input := `BOTPAK	EQU		0x00280000		; bootpackのロード先`
	l := lexer.New(input)
	p := New(l)
	// fmt.Println(p)
	p.ParseProgram()
}
