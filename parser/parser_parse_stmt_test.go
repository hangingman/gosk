package parser

import (
	"fmt"
	"github.com/hangingman/gosk/lexer"
	"github.com/hangingman/gosk/token"
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
	l := lexer.New(input)
	for {
		tok := l.NextToken()
		// fmt.Printf("%s\n", tok)
		if tok.Type == token.EOF {
			break
		}
	}
}
