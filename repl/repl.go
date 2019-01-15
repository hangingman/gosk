package repl

import (
	"bufio"
	"fmt"
	"github.com/hangingman/gosk/eval"
	"github.com/hangingman/gosk/lexer"
	"github.com/hangingman/gosk/object"
	"github.com/hangingman/gosk/parser"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line, logger)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			// error
			continue
		}

		evaluated := eval.Eval(program)
		objArray, ok := evaluated.(*object.ObjectArray)
		if evaluated != nil && ok {
			for _, obj := range *objArray {
				io.WriteString(out, obj.Inspect())
				io.WriteString(out, "\n")
			}
		}
	}
}
