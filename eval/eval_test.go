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
LAHF
LEAVE
LOCK
NOP
OUTSB
OUTSW
OUTSD
POPA
POPAD
POPF
POPFD
PUSHA
PUSHD
PUSHF
RET
RETF
STI
WAIT
`

	tests := []struct {
		Value []byte
	}{
		{[]byte{0x37}},       // AAA
		{[]byte{0x3f}},       // AAS
		{[]byte{0x98}},       // CBW
		{[]byte{0x99}},       // CDQ
		{[]byte{0xf8}},       // CLC
		{[]byte{0xfc}},       // CLD
		{[]byte{0xfa}},       // CLI
		{[]byte{0x0f, 0x06}}, // CLTS
		{[]byte{0xf5}},       // CMC
		{[]byte{0xf8}},       // CPUID
		{[]byte{0x99}},       // CWD
		{[]byte{0x98}},       // CWDE
		{[]byte{0x27}},       // DAA
		{[]byte{0x2f}},       // DAS
		{[]byte{0x9b}},       // WAIT
		{[]byte{0xf4}},       // HLT
		{[]byte{0xce}},       // INCO
		{[]byte{0x6c}},       // INSB
		{[]byte{0x6d}},       // INSD
		{[]byte{0x6d}},       // INSW
		{[]byte{0x0f, 0x08}}, // INVD
		{[]byte{0xcf}},       // IRET
		{[]byte{0xcf}},       // IRETD
		{[]byte{0x9f}},       // LAHF
		{[]byte{0xc9}},       // LEAVE
		{[]byte{0xf0}},       // LOCK
		{[]byte{0x90}},       // NOP
		{[]byte{0x6f}},       // OUTSB
		{[]byte{0x6f}},       // OUTSW
		{[]byte{0x6f}},       // OUTSD
		{[]byte{0x61}},       // POPA
		{[]byte{0x61}},       // POPAD
		{[]byte{0x9d}},       // POPF
		{[]byte{0x9d}},       // POPFD
		{[]byte{0x60}},       // PUSHA
		{[]byte{0x60}},       // PUSHD
		{[]byte{0x9c}},       // PUSHF
		{[]byte{0xc3}},       // RET
		{[]byte{0xcb}},       // RETF
		{[]byte{0xfb}},       // STI
		{[]byte{0x9b}},       // WAIT
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
