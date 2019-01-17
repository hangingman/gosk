package repl

import (
	"bufio"
	"fmt"
	logger "github.com/apsdehal/go-logger"
	"github.com/hangingman/gosk/eval"
	"github.com/hangingman/gosk/lexer"
	"github.com/hangingman/gosk/object"
	"github.com/hangingman/gosk/parser"
	"io"
	"os"
)

const PROMPT = ">> "

var (
	// ロガー
	log, _ = logger.New("repl", 1, os.Stdout)
)

func init() {
	log.SetFormat("[%{module}] [%{level}] %{message}")
	log.SetLogLevel(logger.InfoLevel)
}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
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
