package lexer

import (
	"io/ioutil"
	"fmt"
	"github.com/hangingman/gosk/token"
	"testing"
)


func TestAsmHead(t *testing.T) {
	b, err := ioutil.ReadFile("asmhead.nas")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	l := New(input)
	for {
		tok := l.NextToken()
		fmt.Printf("%s\n", tok)
		if tok.Type == token.EOF {
			break
		}
	}
}
