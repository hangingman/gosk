package eval

import (
	"fmt"
	//"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/lexer"
	"github.com/hangingman/gosk/parser"
	//"github.com/hangingman/gosk/token"
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
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	evaluated := Eval(program)
	assert.Equal(t, "*object.ObjectArray", reflect.TypeOf(evaluated).String())
}
