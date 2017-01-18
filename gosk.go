package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

type Scan struct {
	line     int
	lineHead int
}

func (self *Scan) Init() {
	self.line = 0     // 現在の行
	self.lineHead = 0 // 行の先頭文字位置
}

func (self *Scan) Err(s int) {
	fmt.Printf("\n!!error!!%d\n", s)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:  [--help | --version] source [object/binary] [list]")
		os.Exit(0)
	}
	var buf []byte
	var err error

	buf,err = ioutil.ReadFile(os.Args[1])
	parser := &Parser{Buffer: string(buf)}
	parser.Init()
	parser.s.Init()

	err = parser.Parse()

	if err != nil {
		fmt.Println(err)
	} else {
		parser.Execute()
	}
}
