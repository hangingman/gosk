package eval

import (
	"fmt"
	"github.com/hangingman/gosk/lexer"
	"github.com/hangingman/gosk/object"
	"github.com/hangingman/gosk/parser"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"runtime"
	"testing"
)

func TestEvalSingleOpcode(t *testing.T) {
	input := `AAA
AAS
CBW
CDQ
CLC
CLD
CLI
CLTS
CMC
CPUID
CWD
CWDE
DAA
DAS
WAIT
HLT
INCO
INSB
INSD
INSW
INVD
IRET
IRETD
NOP
POPA
POPF
PUSHA
PUSHD
RET
RETF
STI
WAIT
`

	tests := []struct {
		Value []byte
	}{
		{[]byte{0x37}},
		{[]byte{0x3f}},
		{[]byte{0x98}},
		{[]byte{0x99}},
		{[]byte{0xf8}},
		{[]byte{0xfc}},
		{[]byte{0xfa}},
		{[]byte{0x0f, 0x06}},
		{[]byte{0xf5}},
		{[]byte{0xf8}},
		{[]byte{0x99}},
		{[]byte{0x98}},
		{[]byte{0x27}},
		{[]byte{0x2f}},
		{[]byte{0x9b}},
		{[]byte{0xf4}},
		{[]byte{0xce}},
		{[]byte{0x6c}},
		{[]byte{0x6d}},
		{[]byte{0x6d}},
		{[]byte{0x0f, 0x08}},
		{[]byte{0xcf}},
		{[]byte{0xcf}},
		{[]byte{0x90}},
		{[]byte{0x61}},
		{[]byte{0x9d}},
		{[]byte{0x60}},
		{[]byte{0x60}},
		{[]byte{0xc3}},
		{[]byte{0xcb}},
		{[]byte{0xfb}},
		{[]byte{0x9b}},
	}

	l := lexer.New(input)
	p := parser.New(l)

	// プログラムの解析と評価
	program := p.ParseProgram()
	evaluated := Eval(program)
	// リフレクションで結果をチェック
	assert.Equal(t, "*object.ObjectArray", reflect.TypeOf(evaluated).String())
	// キャストをやる
	objArray, ok := evaluated.(*object.ObjectArray)
	assert.True(t, ok)
	// 結果を１つずつ見てみる
	for i, obj := range *objArray {
		assert.NotEqual(t, obj, nil)
		bin, ok := obj.(*object.Binary)
		assert.True(t, ok)
		assert.Equal(t, tests[i].Value, bin.Value)
	}
}

func TestEvalAsmHead(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	asmheadPath := path.Join(path.Dir(filename), "..", "testdata", "helloos.nas")
	err := os.Chdir("../../")

	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadFile(asmheadPath)
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)

	l := lexer.New(input)
	p := parser.New(l)

	// プログラムの解析と評価
	program := p.ParseProgram()
	evaluated := Eval(program)
	// リフレクションで結果をチェック
	assert.Equal(t, "*object.ObjectArray", reflect.TypeOf(evaluated).String())
	// キャストをやる
	objArray, ok := evaluated.(*object.ObjectArray)
	assert.True(t, ok)
	// 結果を１つずつ見てみる
	for _, obj := range *objArray {
		if obj != nil {
			assert.Equal(t, "*object.Binary", reflect.TypeOf(obj).String())
		}
	}
}
