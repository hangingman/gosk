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

	logger.SetOutput(os.Stdout)
	l := lexer.New(input, logger)
	p := parser.New(l)

	// プログラムの解析と評価
	program := p.ParseProgram()
	evaluated := Eval(program)
	// リフレクションで結果をチェック
	assert.Equal(t, "*object.ObjectArray", reflect.TypeOf(evaluated).String())
	// キャストをやる
	objArray, ok := evaluated.(*object.ObjectArray)
	assert.True(t, ok)
	assert.Equal(t, 27, len(*objArray))
	// 結果を１つずつ見てみる
	for _, obj := range *objArray {
		if obj != nil {
			assert.Equal(t, "*object.Binary", reflect.TypeOf(obj).String())
			// fmt.Printf("%s: %x\n", reflect.TypeOf(obj), obj)
		}
	}
}
